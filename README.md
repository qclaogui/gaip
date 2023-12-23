<div align="center">
  <h1>golang AIP server demo</h1>
</div>
<p align="center">

<a href="https://github.com/qclaogui/gaip/actions/workflows/ci.yml">
  <img src="https://github.com/qclaogui/gaip/actions/workflows/ci.yml/badge.svg">
</a>

<a href="https://goreportcard.com/report/github.com/qclaogui/gaip">
  <img src="https://goreportcard.com/badge/github.com/qclaogui/gaip?v=1" />
</a>

<a href="https://hub.docker.com/r/qclaogui/gaip">
  <img src="https://img.shields.io/docker/pulls/qclaogui/gaip.svg">
</a>

<a href="https://github.com/qclaogui/gaip/blob/master/LICENSE">
  <img src="https://img.shields.io/github/license/qclaogui/gaip.svg" alt="License">
</a>

</p>

Practices for implementing [Google API Improvement Proposals](https://aip.dev/) (AIP) in Go.

Apply knowledge and experience to improve open source

```shell
❯ make help

Usage:
  make <target>

Build
  build                                     Build binary for current OS and place it at ./bin/gaip_$(GOOS)_$(GOARCH)
  build-all                                 Build binaries for Linux and Mac and place them in dist/
  clean                                     Remove artefacts or generated files from previous build

Dependencies
  check-go-mod                              Ensures fresh go.mod and go.sum.
  install-build-deps                        Install dependencies tools

Ent Schema
  ent-gen                                   Regenerate schema
  ent-describe                              Get a description of graph schema
  atlas-lint                                Verifying and linting migrations

Regenerate gRPC Code
  protoc-install                            Install proper protoc version
  protoc-gen                                Regenerate proto by protoc

Testing Lint & Fmt
  test                                      Run tests.
  lint                                      Runs various static analysis against our code.
  fmt                                       Runs fmt code (automatically fix lint errors)

Kubernetes
  cluster                                   Create k3s cluster
  manifests                                 Generates Kubernetes manifests

Release
  prepare-release-candidate                 Create release candidate
  prepare-release                           Create release
  print-version                             Prints the upcoming release number

General
  reference-help                            Generates the reference help documentation.
  help                                      Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/
```

[Automatic Migration planning script](https://entgo.io/docs/versioned/programmatically#2-automatic-migration-planning-script)

```shell
docker run --rm --name atlas-db-dev -d -p 3306:3306 -e MYSQL_DATABASE=dev -e MYSQL_ROOT_PASSWORD=pass mysql:8
```
