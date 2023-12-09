# syntax=docker/dockerfile:1

##
## Build the image
##
FROM golang:1.21.5-bullseye AS builder

# Set destination for COPY
WORKDIR /workspace

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /gaip cmd/main.go

##
## Run the tests in the container
##
FROM builder
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##
FROM gcr.io/distroless/static

# https://github.com/opencontainers/image-spec/blob/main/annotations.md#pre-defined-annotation-keys
LABEL org.opencontainers.image.title="gaip" \
      org.opencontainers.image.source="https://github.com/qclaogui/golang-api-server" \
      org.opencontainers.image.description="Practices for implementing Google API Improvement Proposals (AIP) in Go."

COPY   --from=builder /gaip /usr/local/bin/gaip
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["gaip"]