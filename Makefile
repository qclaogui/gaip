include .bingo/Variables.mk

.DEFAULT_GOAL := help

SWAGGER_UI_VERSION	:=v5.5.0
PROTOC_VERSION		:=24.3

# Download the proper protoc version for Darwin (osx) and Linux.
PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Linux)
	PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip
endif
ifeq ($(UNAME_S), Darwin)
	PROTOC_ZIP=protoc-${PROTOC_VERSION}-osx-universal_binary.zip
endif

PROTOC :=${GOBIN}/protoc-${PROTOC_VERSION}

##@ Build

GOOS             ?= $(shell go env GOOS)
GOARCH           ?= $(shell go env GOARCH)
GOARM            ?= $(shell go env GOARM)
CGO_ENABLED      ?= 0

GO_FILES_TO_FMT  ?= $(shell find . -path ./vendor -prune -o -path ./genproto -prune -o -name '*.go' -print)
# Support gsed on OSX (brew install gnu-sed), falling back to sed. On Linux
# systems gsed won't be installed, so will use sed as expected.
SED ?= $(shell which gsed 2>/dev/null || which sed)

GOPROXY          ?= https://proxy.golang.org
export GOPROXY

GO_ENV := GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) CGO_ENABLED=$(CGO_ENABLED)

VERSION      ?= $(shell ./tools/image-tag)
COMMIT_NO    ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_COMMIT 	 ?= $(if $(shell git status --porcelain --untracked-files=no),${COMMIT_NO}-dirty,${COMMIT_NO})
VPREFIX      := github.com/qclaogui/golang-api-server/pkg/version

GO_LDFLAGS   := -X $(VPREFIX).Version=$(VERSION)                         \
                -X $(VPREFIX).GitCommit=$(GIT_COMMIT)                    \
                -X $(VPREFIX).BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

GO_FLAGS := -ldflags "-s -w $(GO_LDFLAGS)"

.PHONY: build
build: ## Build binary for current OS and place it at ./bin/golang-api-server
	@$(GO_ENV) go build $(GO_FLAGS) -o bin/golang-api-server ./cmd/server

.PHONY: build-all
build-all: ## Build binaries for Linux, Windows and Mac and place them in dist/
	PRE_RELEASE_ID="" $(GORELEASER) --config=.goreleaser.yml --snapshot --skip-publish --clean

.PHONY: clean
clean: ## Remove artefacts or generated files from previous build
	@rm -rf bin dist


##@ Dependencies

.PHONY: go-mod
go-mod: ## go mod download && go mod tidy
	@go mod download
	@go mod tidy
	@go mod verify

.PHONY: check-go-mod
check-go-mod: go-mod ## Ensures fresh go.mod and go.sum.
	@git --no-pager diff --exit-code -- go.sum go.mod vendor/ || { echo ">> There are unstaged changes in go vendoring run 'make go-mod'"; exit 1; }

.PHONY: buf-mod
buf-mod: ## Run buf mod update after adding a dependency to your buf.yaml
	@echo ">> run buf mod update"
	@cd proto/ && $(BUF) mod update

.PHONY: protoc-install
protoc-install:
ifeq ("$(wildcard $(PROTOC))","")
	@cd proto && curl -LO $(PROTOC_URL)$(PROTOC_ZIP)
	@cd proto && unzip -n $(PROTOC_ZIP)
	@cd proto && rm -Rf google/protobuf
	@cd proto && mv -f bin/protoc ${GOBIN}/protoc-${PROTOC_VERSION} && mv -f include/google/protobuf google
	@cd proto && rm -Rf bin include readme.txt $(PROTOC_ZIP)
endif

.PHONY: install-build-deps
install-build-deps: protoc-install ## Install dependencies tools
	$(info ******************** downloading dependencies ********************)
	@echo ">> building bingo and setup dependencies tools"
	@go install github.com/bwplotka/bingo@0568407746a2915ba57f9fa1def47694728b831e


##@ Regenerate gRPC code

.PHONY: buf-gen
buf-gen: ## Regenerate proto by buf https://buf.build/
buf-gen: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2)
	@#rm -Rf genproto third_party/gen
	@cd proto/ && $(BUF) generate \
		--path qclaogui/library/v1/*.proto
	@make swagger-ui
	@make lint

.PHONY: swagger-ui
swagger-ui: ## Generate Swagger UI
	@SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) tools/scripts/generate-swagger-ui.sh

.PHONY: protoc-gen
protoc-gen: ## Regenerate proto by protoc
protoc-gen: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GO_GAPIC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2)
	@rm -Rf genproto third_party/gen
	@mkdir -p genproto third_party/gen/openapiv2
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=genproto \
		--go_opt='module=github.com/qclaogui/golang-api-server/genproto' \
 		proto/qclaogui/routeguide/v1/*.proto \
 		proto/qclaogui/todo/v1/*.proto \
 		proto/qclaogui/project/v1/*.proto \
 		proto/qclaogui/library/v1/*.proto \
 		proto/qclaogui/bookstore/v1alpha1/*.proto

    # plugin protoc-gen-go-grpc
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go-grpc_out=genproto \
		--go-grpc_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--go-grpc_opt='require_unimplemented_servers=false' \
 		proto/qclaogui/routeguide/v1/*.proto \
 		proto/qclaogui/todo/v1/*.proto \
 		proto/qclaogui/project/v1/*.proto \
 		proto/qclaogui/library/v1/*.proto \
 		proto/qclaogui/bookstore/v1alpha1/*.proto

    # plugin protoc-gen-go_gapic
    # https://github.com/googleapis/gapic-generator-go?tab=readme-ov-file#invocation
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/golang-api-server/genproto/bookstore/apiv1alpha1;bookstore' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/bookstore/v1alpha1/bookstore_grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/bookstore/v1alpha1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/golang-api-server/genproto/todo/apiv1;todo' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/todo/v1/todo_grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/todo/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/golang-api-server/genproto/project/apiv1;project' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/project/v1/project_grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/project/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/golang-api-server/genproto/library/apiv1;library' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/library/v1/library_grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/library/v1/*.proto

    # plugin protoc-gen-grpc-gateway
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--grpc-gateway_out=genproto \
		--grpc-gateway_opt='logtostderr=true' \
		--grpc-gateway_opt='module=github.com/qclaogui/golang-api-server/genproto' \
		--grpc-gateway_opt='generate_unbound_methods=true' \
 		proto/qclaogui/todo/v1/*.proto

    # plugin protoc-gen-openapiv2
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
 		--openapiv2_out=third_party/gen/openapiv2 \
 		--openapiv2_opt='logtostderr=true' \
 		--openapiv2_opt='generate_unbound_methods=true' \
 		proto/qclaogui/todo/v1/*.proto

	@make swagger-ui
	@make lint


##@ Testing Lint & fmt

.PHONY: test
test: ## Run tests.
	@$(GO_ENV) go test $(GO_FLAGS) -timeout 10m -count 1 ./...


.PHONY: lint
lint: ## Runs various static analysis against our code.
lint: go-lint goreleaser-lint buf-lint $(COPYRIGHT) fmt
	@$(COPYRIGHT) $(shell go list -f "{{.Dir}}" ./... | grep -iv "genproto/" | xargs -I {} find {} -name "*.go")


.PHONY: fmt
fmt: ## Runs fmt code (automatically fix lint errors)
fmt: fix-lint go-fmt buf-fmt


.PHONY: go-fmt
go-fmt: $(GOIMPORTS) ## Runs gofmt code
	@echo ">> formatting go code"
	@gofmt -s -w $(GO_FILES_TO_FMT)
	@for file in $(GO_FILES_TO_FMT) ; do \
		tools/scripts/goimports.sh "$${file}"; \
	done
	@$(GOIMPORTS) -w $(GO_FILES_TO_FMT)

.PHONY: buf-fmt
buf-fmt: ## examining all of the proto files.
	@echo ">> run buf format"
	@cd proto/ && $(BUF) format -w --exit-code

.PHONY: goreleaser-lint
goreleaser-lint: $(GORELEASER) ## Lint .goreleaser*.yml files.
	@echo ">> run goreleaser check"
	@for config_file in $(shell ls .goreleaser*); do cat $${config_file} > .goreleaser.combined.yml; done
	@$(GORELEASER) check -f .goreleaser.combined.yml || exit 1 && rm .goreleaser.combined.yml

.PHONY: go-lint
go-lint: $(GOLANGCI_LINT) ## examining all of the Go files.
	@echo ">> run golangci-lint"
	@$(GOLANGCI_LINT) run --out-format=github-actions --timeout=15m

.PHONY: buf-lint
buf-lint: $(BUF) buf-fmt ## Lint all of the proto files.
	@echo ">> run buf lint"
	@cd proto/ && $(BUF) lint

.PHONY: fix-lint
fix-lint: $(GOLANGCI_LINT) ## fix lint issue of the Go files
	@echo ">> fix lint issue of the Go files"
	@$(GOLANGCI_LINT) run --fix


##@ Release

.PHONY: prepare-release-candidate
prepare-release-candidate: ## Create release candidate
	tools/scripts/tag-release-candidate.sh

.PHONY: prepare-release
prepare-release: ## Create release
	tools/scripts/tag-release.sh

.PHONY: print-version
print-version: ## Prints the upcoming release number
	@go run pkg/version/generate/release_generate.go print-version

.PHONY: manifests
manifests: $(KUSTOMIZE) ## Generates the k8s manifests
	@$(KUSTOMIZE) build deploy/overlays/dev > deploy/overlays/dev/k8s-all-in-one.yaml
	@$(KUSTOMIZE) build deploy/overlays/prod > deploy/overlays/prod/k8s-all-in-one.yaml


##@ General

.PHONY: help
help:  ## Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/
ifeq ($(OS),Windows_NT)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  %-40s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
else
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
endif
