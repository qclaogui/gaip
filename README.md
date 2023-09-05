<div align="center">
  <h1>golang api server demo</h1>
</div>
<p align="center">

<a href="https://github.com/qclaogui/golang-api-server/actions/workflows/ci.yml">
  <img src="https://github.com/qclaogui/golang-api-server/actions/workflows/ci.yml/badge.svg">
</a>

<a href="https://goreportcard.com/report/github.com/qclaogui/golang-api-server">
  <img src="https://goreportcard.com/badge/github.com/qclaogui/golang-api-server?v=1" />
</a>

<a href="https://hub.docker.com/r/qclaogui/golang-api-server">
  <img src="https://img.shields.io/docker/pulls/qclaogui/golang-api-server.svg">
</a>

<a href="https://github.com/qclaogui/golang-api-server/blob/master/LICENSE">
  <img src="https://img.shields.io/github/license/qclaogui/golang-api-server.svg" alt="License">
</a>

</p>

Share knowledge and help others.

```shell
❯ make help

Usage:
  make <target>

Build

Regenerate gRPC code
  buf-gen                                   Regenerate proto by buf https://buf.build/
  swagger-ui                                Generate Swagger UI
  protoc-gen                                Regenerate proto by protoc

Dependencies
  go-mod                                    go mod download && go mod tidy
  check-go-mod                              Ensures fresh go.mod and go.sum.
  install-build-deps                        Install dependencies tools
  build                                     Build golang-api-server binary for current OS and place it at ./bin/golang-api-server
  build-all                                 Build binaries for Linux, Windows and Mac and place them in dist/
  clean                                     Remove artefacts or generated files from previous build

Testing Lint & fmt
  fmt                                       Runs fmt code. (go-fmt buf-fmt)
  go-fmt                                    Runs gofmt code
  buf-mod                                   Run buf mod update after adding a dependency to your buf.yaml
  buf-fmt                                   examining all of the proto files.
  lint                                      Runs various static analysis against our code.
  goreleaser-lint                           Lint .goreleaser*.yml files.
  go-lint                                   examining all of the Go files.
  buf-lint                                  Lint all of the proto files.
  fix-lint                                  fix lint issue of the Go files
  test                                      Run tests.

Release
  prepare-release-candidate                 Create release candidate
  prepare-release                           Create release
  print-version                             Prints the upcoming release number
  manifests                                 Generates the k8s manifests

General
  help                                      Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/

```
