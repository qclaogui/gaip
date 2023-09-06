# syntax=docker/dockerfile:1

##
## Build the image
##
FROM golang:1.21.1-bullseye AS builder
LABEL maintainer="qclaogui@gmail.com"

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main cmd/main.go

##
## Run the tests in the container
##
FROM builder
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##
FROM gcr.io/distroless/static

COPY --from=builder /main /usr/local/bin/main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["main"]