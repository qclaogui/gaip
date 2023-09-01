include .bingo/Variables.mk

.DEFAULT_GOAL := help

##@ Build

#
# Environment variables:
#

GOOS             ?= $(shell go env GOOS)
GOARCH           ?= $(shell go env GOARCH)
GOARM            ?= $(shell go env GOARM)
CGO_ENABLED      ?= 0

GO_FILES_TO_FMT  ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)
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

##@ Regenerate gRPC code

.PHONY: buf/gen
buf/gen: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) ## buf regenerate gRPC code
	@rm -Rf api/openapiv2/gen/ api/gen
	cd api/ && $(BUF) generate

.PHONY: protoc/gen
protoc/gen: $(PROTOC_GEN_GO) ## protoc regenerate gRPC code
	@protoc -I api \
		--go_out api/gen/proto \
		--go_opt paths=source_relative \
		--go_grpc_out api/gen/proto \
		--go_grpc_opt paths=source_relative \
		--go_grpc_opt require_unimplemented_servers=false \
		api/todo/v1/todo_service.proto  \
		api/routeguide/v1/route_guide.proto

# grpc-gateway
# https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/
	@protoc -I api \
		--grpc-gateway_out api/gen/proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		api/todo/v1/todo_service.proto

# # grpc-gateway
# # https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/
# 	@protoc -I api \
# 		--grpc-gateway_out api/gen/proto \
# 		--grpc-gateway_opt logtostderr=true \
# 		--grpc-gateway_opt paths=source_relative \
# 		--grpc-gateway_opt generate_unbound_methods=true \
# 		--grpc-gateway_opt standalone=true \
# 		--grpc-gateway_opt grpc_api_configuration=path/to/your_service.yaml \
# 		api/todo/v1/todo_service.proto

# # OpenAPI
# 	@protoc -I api \
# 		--openapiv2_out api/gen/proto \
# 		--openapiv2_opt logtostderr=true \
# 		--openapiv2_opt grpc_api_configuration=path/to/your_service.yaml \
# 		--openapiv2_opt openapi_configuration=path/to/your_service_swagger.yaml \
# 		api/todo/v1/todo_service.proto

##@ Dependencies

.PHONY: go/mod
go/mod: ## Ensures fresh go.mod and go.sum.
	@go mod download
	@go mod tidy
	@go mod verify

.PHONY: check/mod
check/mod: go/mod
	@git --no-pager diff --exit-code -- go.sum go.mod vendor/ || { echo ">> There are unstaged changes in go vendoring run 'make go/mod'"; exit 1; }

.PHONY: install-build-deps
install-build-deps: ## Install dependencies tools
	$(info ******************** downloading dependencies ********************)
	@echo ">> building bingo and setup dependencies tools"
	@go install github.com/bwplotka/bingo@0568407746a2915ba57f9fa1def47694728b831e

.PHONY: build
build: ## Build golang-api-server binary for current OS and place it at ./bin/api-server
	@$(GO_ENV) go build $(GO_FLAGS) -o bin/golang-api-server ./cmd

.PHONY: build/all
build/all: ## Build binaries for Linux, Windows and Mac and place them in dist/
	PRE_RELEASE_ID="" $(GORELEASER) --config=.goreleaser.yml --snapshot --skip-publish --clean

.PHONY: clean
clean: ## Remove artefacts or generated files from previous build
	rm -rf bin dist

##@ Testing Lint & fmt

.PHONY: fmt
fmt: ## Runs fmt code.
fmt: go/fmt buf/fmt
	$(info ******************** done ********************)

.PHONY: go/fmt
go/fmt: $(GOIMPORTS)
	@echo ">> formatting go code"
	@gofmt -s -w $(GO_FILES_TO_FMT)
	@for file in $(GO_FILES_TO_FMT) ; do \
		tools/scripts/goimports.sh "$${file}"; \
	done
	@$(GOIMPORTS) -w $(GO_FILES_TO_FMT)

.PHONY: buf/fmt
buf/fmt: ## examining all of the proto files.
	@echo ">> run buf format"
	@cd api/ && $(BUF) format -w --exit-code

.PHONY: lint
lint: ## Runs various static analysis against our code.
lint: go/lint goreleaser/lint buf/lint $(COPYRIGHT)
	@$(COPYRIGHT) $(shell go list -f "{{.Dir}}" ./... | xargs -I {} find {} -name "*.go")
	$(info ******************** done ********************)

.PHONY: goreleaser/lint
goreleaser/lint: $(GORELEASER) ## examining all of the Go files.
	@echo ">> run goreleaser check"
	@for config_file in $(shell ls .goreleaser*); do cat $${config_file} > .goreleaser.combined.yml; done
	$(GORELEASER) check -f .goreleaser.combined.yml || exit 1 && rm .goreleaser.combined.yml

.PHONY: go/lint
go/lint: $(GOLANGCI_LINT) ## examining all of the Go files.
	@echo ">> run golangci-lint"
	$(GOLANGCI_LINT) run --out-format=github-actions --timeout=15m

.PHONY: buf/lint
buf/lint: $(BUF) buf/fmt ## examining all of the proto files.
	@echo ">> run buf lint"
	@cd api/ && $(BUF) lint

.PHONY: fix/lint
fix/lint: $(GOLANGCI_LINT) ## examining all of the Go files.
	@echo ">> run golangci-lint fix"
	$(GOLANGCI_LINT) run --fix

.PHONY: test
test: ## Run tests.
	@$(GO_ENV) go test $(GO_FLAGS) -timeout 10m -count 1 ./...

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
