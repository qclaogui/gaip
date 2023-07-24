.DEFAULT_GOAL := help

##@ Build

#
# Environment variables:
#

GOOS             ?= $(shell go env GOOS)
GOARCH           ?= $(shell go env GOARCH)
GOARM            ?= $(shell go env GOARM)
CGO_ENABLED      ?= 1

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
build: ## Build golang-api-server binary for current OS and place it at ./bin/api-server
	$(GO_ENV) go build $(GO_FLAGS) -o bin/golang-api-server ./cmd

.PHONY: build-all
build-all: ## Build binaries for Linux, Windows and Mac and place them in dist/
	PRE_RELEASE_ID="" goreleaser --config=.goreleaser.yml --snapshot --skip-publish --clean

.PHONY: clean
clean: ## Remove artefacts or generated files from previous build
	rm -rf bin dist

##@ Testing & CI

.PHONY: lint
lint: ## Run linter over the codebase
	golangci-lint run --out-format=github-actions --timeout=15m
	@for config_file in $(shell ls .goreleaser*); do goreleaser check -f $${config_file} || exit 1; done

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	$(GO_ENV) go test $(GO_FLAGS) -timeout 10m -race -count 1 ./...


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
manifests: ## Generates the k8s manifests
	kustomize build deploy/overlays/dev > deploy/overlays/dev/k8s-all-in-one.yaml
	kustomize build deploy/overlays/prod > deploy/overlays/prod/k8s-all-in-one.yaml

##@ General

.PHONY: help
help:  ## Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/
ifeq ($(OS),Windows_NT)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  %-40s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
else
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
endif
