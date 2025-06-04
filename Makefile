include .bingo/Variables.mk

.DEFAULT_GOAL := help

SWAGGER_UI_VERSION	:=v5.22.0
PROTOC_VERSION		:=31.1

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
CGO_ENABLED      ?= 1
RELEASE_BUILD    ?= 0

GO_FILES_TO_FMT  ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)

# Support gsed on OSX (brew install gnu-sed), falling back to sed. On Linux
# systems gsed won't be installed, so will use sed as expected.
SED ?= $(shell which gsed 2>/dev/null || which sed)

GOPROXY          ?= https://proxy.golang.org
export GOPROXY

GO_ENV := GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) CGO_ENABLED=$(CGO_ENABLED)

VERSION 	?= $(shell ./tools/image-tag)
COMMIT_NO 	?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_COMMIT 	?= $(if $(shell git status --porcelain --untracked-files=no),${COMMIT_NO}-dirty,${COMMIT_NO})
VPREFIX 	:= github.com/qclaogui/gaip/pkg/version

GO_LDFLAGS   := -X $(VPREFIX).Version=$(VERSION)                         \
                -X $(VPREFIX).GitCommit=$(GIT_COMMIT)                    \
                -X $(VPREFIX).BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

DEFAULT_FLAGS	:= $(GO_FLAGS)

ifeq ($(RELEASE_BUILD),1)
GO_FLAGS	:= $(DEFAULT_FLAGS) -ldflags "-s -w $(GO_LDFLAGS)"
else
GO_FLAGS	:= $(DEFAULT_FLAGS) -ldflags "$(GO_LDFLAGS)"
endif

.PHONY: build
build: ## Build binary for current OS and place it at ./bin/gaip_$(GOOS)_$(GOARCH)
	@$(GO_ENV) go build $(GO_FLAGS) -o bin/gaip_$(GOOS)_$(GOARCH) ./cmd/gaip

.PHONY: build-all
build-all: ## Build binaries for Linux and Mac and place them in dist/
	@cat .github/.goreleaser.yml .github/.goreleaser.docker.yml > .github/.goreleaser.combined.yml
	RELEASE_BUILD=$(RELEASE_BUILD) PRE_RELEASE_ID="" $(GORELEASER) --config=.github/.goreleaser.combined.yml --snapshot --skip=publish --clean
	@rm .github/.goreleaser.combined.yml

.PHONY: clean
clean: ## Remove artefacts or generated files from previous build
	@rm -rf bin dist


##@ Dependencies

go-mod:
	@go mod download
	@go mod tidy
	@go mod verify

check-go-mod: go-mod ## Ensures fresh go.mod and go.sum.
	@git --no-pager diff --exit-code -- go.sum go.mod vendor/ || { echo ">> There are unstaged changes in go vendoring run 'make go-mod'"; exit 1; }

# .PHONY: buf-mod
# buf-mod: ## Run buf mod update after adding a dependency to your buf.yaml
# 	@echo ">> run buf mod update"
# 	@cd proto/ && $(BUF) mod update

install-build-deps: ## Install dependencies tools
	$(info ******************** downloading dependencies ********************)
	@echo ">> building bingo and setup dependencies tools"
	@go install github.com/bwplotka/bingo@v0.9.0


##@ Ent Schema

# Generate the schema under internal/ent/schema/ directory
.PHONY: ent-new
ent-new: $(ENT)
	@$(ENT) new --target=internal/ent/schema \
			Todo

ent-gen: ## Regenerate schema
	@go generate ./internal/ent

ent-describe: $(ENT) ## Get a description of graph schema
	@$(ENT) describe ./internal/ent/schema

atlas-lint: $(ATLAS) ## Verifying and linting migrations
	@$(ATLAS) migrate lint \
      --dir "file://migrations" \
      --dev-url "docker://mysql/8/test" \
      --latest 1

# Generating Versioned Migration Files
.PHONY: atlas-diff
atlas-diff: $(ATLAS)
	@$(ATLAS) migrate diff migration_name \
      --dir "file://migrations" \
      --to "ent://internal/ent/schema" \
      --dev-url "docker://mysql/8/ent"

# Apply generated migration files onto the database
.PHONY: atlas-apply
atlas-apply: $(ATLAS)
	@$(ATLAS) migrate apply \
      --dir="file://migrations" \
      --url="mysql://root:pass@localhost:3306/example"

##@ Regenerate gRPC Code

protoc-install: ## Install proper protoc version
ifeq ("$(wildcard $(PROTOC))","")
	@cd proto && curl -LO $(PROTOC_URL)$(PROTOC_ZIP)
	@cd proto && unzip -n $(PROTOC_ZIP)
	@cd proto && rm -Rf google/protobuf
	@cd proto && mv -f bin/protoc ${GOBIN}/protoc-${PROTOC_VERSION} && mv -f include/google/protobuf google
	@cd proto && rm -Rf bin include readme.txt $(PROTOC_ZIP)
endif
	@echo ">> (re)installing protobuf and proper protoc version"

# buf-gen: ## Regenerate proto by buf https://buf.build/
# buf-gen: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2)
# #	@rm -Rf genproto thirdparty/gen
# 	@cd proto/ && $(BUF) generate \
# 		--path qclaogui/library/v1
# 	@make swagger-ui
# 	@make lint

# Generate Swagger UI
.PHONY: swagger-ui
swagger-ui:
	@SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) tools/scripts/generate-swagger-ui.sh

protoc-gen: ## Regenerate proto by protoc
protoc-gen: protoc-install $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GO_GAPIC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_OPENAPIV2) $(PROTOC_GEN_GO_REST_HANDLER)
	@rm -Rf genproto thirdparty/gen
	@mkdir -p genproto thirdparty/gen/openapiv2
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=genproto \
		--go_opt='module=github.com/qclaogui/gaip/genproto' \
 		proto/qclaogui/bookstore/v1alpha1/*.proto \
 		proto/qclaogui/generativelanguage/v1/*.proto \
 		proto/qclaogui/generativelanguage/v1beta/*.proto \
 		proto/qclaogui/aiplatform/v1beta1/*.proto \
 		proto/qclaogui/library/v1/*.proto \
 		proto/qclaogui/project/v1/*.proto \
 		proto/qclaogui/showcase/v1beta1/*.proto \
 		proto/qclaogui/task/v1/*.proto \
 		proto/qclaogui/todo/v1/*.proto

    # plugin protoc-gen-go-grpc
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go-grpc_out=genproto \
		--go-grpc_opt='module=github.com/qclaogui/gaip/genproto' \
		--go-grpc_opt='require_unimplemented_servers=false' \
 		proto/qclaogui/bookstore/v1alpha1/*.proto \
 		proto/qclaogui/generativelanguage/v1/*.proto \
 		proto/qclaogui/generativelanguage/v1beta/*.proto \
 		proto/qclaogui/aiplatform/v1beta1/*.proto \
 		proto/qclaogui/library/v1/*.proto \
 		proto/qclaogui/project/v1/*.proto \
 		proto/qclaogui/showcase/v1beta1/*.proto \
 		proto/qclaogui/task/v1/*.proto \
 		proto/qclaogui/todo/v1/*.proto

    # plugin protoc-gen-go_rest_handler
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_rest_handler=$(PROTOC_GEN_GO_REST_HANDLER) \
		--go_rest_handler_out=genproto \
		--go_rest_handler_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_rest_handler_opt='api-service-config=proto/qclaogui/library/v1/library_v1.yaml' \
		proto/qclaogui/library/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_rest_handler=$(PROTOC_GEN_GO_REST_HANDLER) \
		--go_rest_handler_out=genproto \
		--go_rest_handler_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_rest_handler_opt='api-service-config=proto/qclaogui/showcase/v1beta1/showcase_v1beta1.yaml' \
 		proto/qclaogui/showcase/v1beta1/*.proto

#	@$(PROTOC) --proto_path=proto \
#		--plugin=protoc-gen-go_rest_handler=$(PROTOC_GEN_GO_REST_HANDLER) \
#		--go_rest_handler_out=genproto \
#		--go_rest_handler_opt='module=github.com/qclaogui/gaip/genproto' \
#		--go_rest_handler_opt='api-service-config=proto/qclaogui/generativelanguage/v1beta/generativelanguage_v1beta.yaml' \
# 		proto/qclaogui/generativelanguage/v1beta/*.proto


    # plugin protoc-gen-go_gapic
    # https://github.com/googleapis/gapic-generator-go?tab=readme-ov-file#invocation

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/generativelanguage/apiv1;generativelanguage' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/generativelanguage/v1/grpc_service_config.json' \
		--go_gapic_opt='api-service-config=proto/qclaogui/generativelanguage/v1/generativelanguage_v1.yaml' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/generativelanguage/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta;generativelanguage' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/generativelanguage/v1beta/grpc_service_config.json' \
		--go_gapic_opt='api-service-config=proto/qclaogui/generativelanguage/v1beta/generativelanguage_v1beta.yaml' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/generativelanguage/v1beta/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1;aiplatform' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/aiplatform/v1beta1/grpc_service_config.json' \
		--go_gapic_opt='api-service-config=proto/qclaogui/aiplatform/v1beta1/aiplatform_v1beta1.yaml' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/aiplatform/v1beta1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/library/apiv1;library' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/library/v1/grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/library/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/project/apiv1;project' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/project/v1/grpc_service_config.json' \
		--go_gapic_opt='api-service-config=proto/qclaogui/project/v1/project_v1.yaml' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/project/v1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/showcase/apiv1beta1;showcase' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/showcase/v1beta1/grpc_service_config.json' \
		--go_gapic_opt='api-service-config=proto/qclaogui/showcase/v1beta1/showcase_v1beta1.yaml' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/showcase/v1beta1/*.proto

	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-go_gapic=$(PROTOC_GEN_GO_GAPIC) \
		--go_gapic_out=genproto \
		--go_gapic_opt='go-gapic-package=github.com/qclaogui/gaip/genproto/task/apiv1;task' \
		--go_gapic_opt='metadata=false' \
		--go_gapic_opt='omit-snippets' \
		--go_gapic_opt='module=github.com/qclaogui/gaip/genproto' \
		--go_gapic_opt='grpc-service-config=proto/qclaogui/task/v1/grpc_service_config.json' \
		--go_gapic_opt='release-level=alpha' \
		--go_gapic_opt='transport=grpc+rest' \
		--go_gapic_opt='rest-numeric-enums=true' \
 		proto/qclaogui/task/v1/*.proto


    # plugin protoc-gen-grpc-gateway
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--grpc-gateway_out=genproto \
		--grpc-gateway_opt='logtostderr=true' \
		--grpc-gateway_opt='module=github.com/qclaogui/gaip/genproto' \
		--grpc-gateway_opt='grpc_api_configuration=proto/qclaogui/todo/v1/api_config_http.yaml' \
 		proto/qclaogui/todo/v1/*.proto

    # plugin protoc-gen-openapiv2
	@$(PROTOC) --proto_path=proto \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
 		--openapiv2_out=thirdparty/gen/openapiv2 \
 		--openapiv2_opt='logtostderr=true' \
 		--openapiv2_opt='grpc_api_configuration=proto/qclaogui/todo/v1/api_config_http.yaml' \
 		proto/qclaogui/todo/v1/*.proto

	@make swagger-ui
	@make fmt
	@make lint
#


# For the following protoveneer-gen line to work, install the protoveener tool:
#    git clone https://github.com/googleapis/google-cloud-go
#    cd google-cloud-go
#    go install ./internal/protoveneer/cmd/protoveneer
#
protoveneer-gen: $(PROTOVENEER)
	@$(PROTOVENEER) -license LICENSE -outdir vertexai/genai vertexai/genai/protoveneer.yaml genproto/aiplatform/apiv1beta1/aiplatformpb

##@ Testing Lint & Fmt

test: ## Run tests.
	@$(GO_ENV) go test $(GO_FLAGS) -timeout 20m -count 1 ./...

integration-tests: ## Run all integration tests.
integration-tests:
	go test -tags=requires_docker -timeout 20m ./integration/tests/...


lint: ## Runs various static analysis against our code.
lint: go-lint goreleaser-lint buf-lint $(COPYRIGHT) fmt
	@$(COPYRIGHT) $(shell go list -f "{{.Dir}}" ./... | grep -iv "genproto/" | xargs -I {} find {} -name "*.go")

.PHONY: fmt
fmt: ## Runs fmt code (automatically fix lint errors)
fmt: fix-lint go-fmt buf-fmt

.PHONY: go-fmt
go-fmt: $(GOIMPORTS) $(GOFUMPT)
	@echo ">> formatting go code"
	@$(GOFUMPT) -w $(GO_FILES_TO_FMT)
	@$(GOIMPORTS) -w $(GO_FILES_TO_FMT)

buf-fmt: $(BUF)
	@echo ">> run buf format"
	@cd proto/ && $(BUF) format -w --exit-code

# Lint .goreleaser*.yml files.
.PHONY: goreleaser-lint
goreleaser-lint: $(GORELEASER)
	@echo ">> run goreleaser check"
	@for config_file in $(shell ls .github/.goreleaser*); do cat $${config_file} > .github/.goreleaser.combined.yml; done
	@$(GORELEASER) check -f .github/.goreleaser.combined.yml || exit 1 && rm .github/.goreleaser.combined.yml

go-lint: $(GOLANGCI_LINT)
	@echo ">> run golangci-lint"
	@$(GOLANGCI_LINT) run --timeout=15m

buf-lint: $(BUF) buf-fmt
	@echo ">> run buf lint"
	@cd proto/ && $(BUF) lint

api-linter: $(API_LINTER) buf-fmt
	@echo ">> run api-linter lint"
	@cd proto/ && $(API_LINTER) \
	qclaogui/project/v1/echo_service.proto \
	qclaogui/project/v1/project_service.proto \
	--set-exit-status

fix-lint: $(GOLANGCI_LINT)
	@echo ">> fix lint issue of the Go files"
	@$(GOLANGCI_LINT) run --fix



##@ Kubernetes

cluster: ## Create k3s cluster
	k3d cluster create k3s-gaip --config deploy/k3d-k3s-config.yaml
#	k3d image import -c k3s-gaip qclaogui/gaip:latest

.PHONY: manifests
manifests: $(notdir $(wildcard deploy/kustomize/overlays/*)) ## Generates Kubernetes manifests

%: ## Generates overlays manifests
	$(info ******************** generates $@ manifests ********************)
	@$(KUSTOMIZE) build --enable-helm deploy/kustomize/overlays/$@ > deploy/kustomize/overlays/$@/manifests/k8s-all-in-one.yaml


generate-helm-docs: # Docs generated by https://github.com/norwoodj/helm-docs
	cd deploy/helm/charts/gaip && $(HELM_DOCS)

generate-helm-tests:
	bash ./deploy/helm/scripts/rebuild-tests.sh

##@ Release

prepare-release-candidate: ## Create release candidate
	tools/scripts/tag-release-candidate.sh

prepare-release: ## Create release
	tools/scripts/tag-release.sh

print-version: ## Prints the upcoming release number
	@go run pkg/version/generate/release_generate.go print-version


##@ General

reference-help: ## Generates the reference help documentation.
reference-help: build
	@(./bin/gaip_$(GOOS)_$(GOARCH) -help || true) > cmd/gaip/help.txt.tmpl
	@(./bin/gaip_$(GOOS)_$(GOARCH) -help-all || true) > cmd/gaip/help-all.txt.tmpl

skaffold-fix: $(SKAFFOLD) ## Update "skaffold.yaml" in the current folder in-place
	@$(SKAFFOLD) fix --overwrite

help:  ## Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/
ifeq ($(OS),Windows_NT)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  %-40s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
else
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
endif
