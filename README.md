<div align="center">
  <h1>golang AIP server demo</h1>
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

Practices for implementing [Google API Improvement Proposals](https://aip.dev/) (AIP) in Go.

Share knowledge and help others.

```shell
❯ make help

Usage:
  make <target>

Build
  build                                     Build binary for current OS and place it at ./bin/gaip
  build-all                                 Build binaries for Linux, Windows and Mac and place them in dist/
  clean                                     Remove artefacts or generated files from previous build

Dependencies
  go-mod                                    go mod download && go mod tidy
  check-go-mod                              Ensures fresh go.mod and go.sum.
  buf-mod                                   Run buf mod update after adding a dependency to your buf.yaml
  install-build-deps                        Install dependencies tools

Generate the schema under internal/ent/schema/ directory
  ent-new                                   Get a description of graph schema
  ent-gen                                   Regenerate schema
  ent-describe                              Get a description of graph schema
  atlas-lint                                Verifying and linting migrations
  atlas-diff                                Generating Versioned Migration Files
  atlas-apply                               Apply generated migration files onto the database

Regenerate gRPC code
  buf-gen                                   Regenerate proto by buf https://buf.build/
  swagger-ui                                Generate Swagger UI
  protoc-gen                                Regenerate proto by protoc

Testing Lint & fmt
  test                                      Run tests.
  lint                                      Runs various static analysis against our code.
  fmt                                       Runs fmt code (automatically fix lint errors)
  go-fmt                                    Runs gofmt code
  buf-fmt                                   examining all of the proto files.
  goreleaser-lint                           Lint .goreleaser*.yml files.
  go-lint                                   examining all of the Go files.
  buf-lint                                  Lint all of the proto files.
  api-linter                                Lint all of the proto files.
  fix-lint                                  fix lint issue of the Go files

Release
  prepare-release-candidate                 Create release candidate
  prepare-release                           Create release
  print-version                             Prints the upcoming release number
  manifests                                 Generates the k8s manifests

General
  help                                      Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/

```

[Automatic Migration planning script](https://entgo.io/docs/versioned/programmatically#2-automatic-migration-planning-script)

```shell
docker run --rm --name atlas-db-dev -d -p 3306:3306 -e MYSQL_DATABASE=dev -e MYSQL_ROOT_PASSWORD=pass mysql:8
```
