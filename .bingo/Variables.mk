# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.9. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for api-linter variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(API_LINTER)
#	@echo "Running api-linter"
#	@$(API_LINTER) <flags/args..>
#
API_LINTER := $(GOBIN)/api-linter-v1.69.2
$(API_LINTER): $(BINGO_DIR)/api-linter.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/api-linter-v1.69.2"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=api-linter.mod -o=$(GOBIN)/api-linter-v1.69.2 "github.com/googleapis/api-linter/cmd/api-linter"

ATLAS := $(GOBIN)/atlas-v0.13.1
$(ATLAS): $(BINGO_DIR)/atlas.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/atlas-v0.13.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=atlas.mod -o=$(GOBIN)/atlas-v0.13.1 "ariga.io/atlas/cmd/atlas"

BINGO := $(GOBIN)/bingo-v0.9.0
$(BINGO): $(BINGO_DIR)/bingo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/bingo-v0.9.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=bingo.mod -o=$(GOBIN)/bingo-v0.9.0 "github.com/bwplotka/bingo"

BUF := $(GOBIN)/buf-v1.54.0
$(BUF): $(BINGO_DIR)/buf.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/buf-v1.54.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=buf.mod -o=$(GOBIN)/buf-v1.54.0 "github.com/bufbuild/buf/cmd/buf"

COPYRIGHT := $(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60
$(COPYRIGHT): $(BINGO_DIR)/copyright.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=copyright.mod -o=$(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60 "github.com/efficientgo/tools/copyright"

ENT := $(GOBIN)/ent-v0.14.4
$(ENT): $(BINGO_DIR)/ent.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/ent-v0.14.4"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=ent.mod -o=$(GOBIN)/ent-v0.14.4 "entgo.io/ent/cmd/ent"

GOFUMPT := $(GOBIN)/gofumpt-v0.8.0
$(GOFUMPT): $(BINGO_DIR)/gofumpt.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gofumpt-v0.8.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=gofumpt.mod -o=$(GOBIN)/gofumpt-v0.8.0 "mvdan.cc/gofumpt"

GOIMPORTS := $(GOBIN)/goimports-v0.33.0
$(GOIMPORTS): $(BINGO_DIR)/goimports.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goimports-v0.33.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=goimports.mod -o=$(GOBIN)/goimports-v0.33.0 "golang.org/x/tools/cmd/goimports"

GOLANGCI_LINT := $(GOBIN)/golangci-lint-v2.1.6
$(GOLANGCI_LINT): $(BINGO_DIR)/golangci-lint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golangci-lint-v2.1.6"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=golangci-lint.mod -o=$(GOBIN)/golangci-lint-v2.1.6 "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"

GORELEASER := $(GOBIN)/goreleaser-v1.26.2
$(GORELEASER): $(BINGO_DIR)/goreleaser.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goreleaser-v1.26.2"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=goreleaser.mod -o=$(GOBIN)/goreleaser-v1.26.2 "github.com/goreleaser/goreleaser"

HELM_DOCS := $(GOBIN)/helm-docs-v1.14.2
$(HELM_DOCS): $(BINGO_DIR)/helm-docs.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/helm-docs-v1.14.2"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=helm-docs.mod -o=$(GOBIN)/helm-docs-v1.14.2 "github.com/norwoodj/helm-docs/cmd/helm-docs"

KUSTOMIZE := $(GOBIN)/kustomize-v5.6.0
$(KUSTOMIZE): $(BINGO_DIR)/kustomize.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/kustomize-v5.6.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=kustomize.mod -o=$(GOBIN)/kustomize-v5.6.0 "sigs.k8s.io/kustomize/kustomize/v5"

PROTOC_GEN_GO_GRPC := $(GOBIN)/protoc-gen-go-grpc-v1.5.1
$(PROTOC_GEN_GO_GRPC): $(BINGO_DIR)/protoc-gen-go-grpc.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go-grpc-v1.5.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go-grpc.mod -o=$(GOBIN)/protoc-gen-go-grpc-v1.5.1 "google.golang.org/grpc/cmd/protoc-gen-go-grpc"

PROTOC_GEN_GO := $(GOBIN)/protoc-gen-go-v1.36.6
$(PROTOC_GEN_GO): $(BINGO_DIR)/protoc-gen-go.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go-v1.36.6"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go.mod -o=$(GOBIN)/protoc-gen-go-v1.36.6 "google.golang.org/protobuf/cmd/protoc-gen-go"

PROTOC_GEN_GO_GAPIC := $(GOBIN)/protoc-gen-go_gapic-v0.53.1
$(PROTOC_GEN_GO_GAPIC): $(BINGO_DIR)/protoc-gen-go_gapic.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go_gapic-v0.53.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go_gapic.mod -o=$(GOBIN)/protoc-gen-go_gapic-v0.53.1 "github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic"

PROTOC_GEN_GO_REST_HANDLER := $(GOBIN)/protoc-gen-go_rest_handler-v0.0.0-20240516080558-0052bb7d2276
$(PROTOC_GEN_GO_REST_HANDLER): $(BINGO_DIR)/protoc-gen-go_rest_handler.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go_rest_handler-v0.0.0-20240516080558-0052bb7d2276"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go_rest_handler.mod -o=$(GOBIN)/protoc-gen-go_rest_handler-v0.0.0-20240516080558-0052bb7d2276 "github.com/qclaogui/gaip/cmd/protoc-gen-go_rest_handler"

PROTOC_GEN_GRPC_GATEWAY := $(GOBIN)/protoc-gen-grpc-gateway-v2.26.3
$(PROTOC_GEN_GRPC_GATEWAY): $(BINGO_DIR)/protoc-gen-grpc-gateway.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-grpc-gateway-v2.26.3"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-grpc-gateway.mod -o=$(GOBIN)/protoc-gen-grpc-gateway-v2.26.3 "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"

PROTOC_GEN_OPENAPIV2 := $(GOBIN)/protoc-gen-openapiv2-v2.26.3
$(PROTOC_GEN_OPENAPIV2): $(BINGO_DIR)/protoc-gen-openapiv2.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-openapiv2-v2.26.3"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-openapiv2.mod -o=$(GOBIN)/protoc-gen-openapiv2-v2.26.3 "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"

PROTOVENEER := $(GOBIN)/protoveneer-v0.0.0-20250529072452-215f1d052cf0
$(PROTOVENEER): $(BINGO_DIR)/protoveneer.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoveneer-v0.0.0-20250529072452-215f1d052cf0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoveneer.mod -o=$(GOBIN)/protoveneer-v0.0.0-20250529072452-215f1d052cf0 "github.com/qclaogui/gaip/cmd/protoveneer"

SKAFFOLD := $(GOBIN)/skaffold-v2.16.0
$(SKAFFOLD): $(BINGO_DIR)/skaffold.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/skaffold-v2.16.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=skaffold.mod -o=$(GOBIN)/skaffold-v2.16.0 "github.com/GoogleContainerTools/skaffold/v2/cmd/skaffold"

