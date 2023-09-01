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
❯ tree -I 'vendor|docs|tools|deploy|cmd' -L 3
.
├── COPYRIGHT
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── SECURITY.md
├── api
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── gen
│   │   └── proto
│   ├── routeguide
│   │   ├── v1
│   │   └── v1beta1
│   └── todo
│       ├── v1
│       └── v1alpha
├── ci
│   └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── protocol
│   │   ├── grpc
│   │   └── rest
│   ├── service
│   │   ├── routeguide
│   │   └── todo
│   └── version
│       ├── generate
│       ├── release.go
│       └── version.go
├── skaffold.env
└── skaffold.yaml
```
