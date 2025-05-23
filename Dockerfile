# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /workspace

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/gaip_linux_amd64 cmd/main.go

##
## Deploy the application binary into a lean image
##
FROM gcr.io/distroless/static

LABEL  org.opencontainers.image.title="gaip" \
       org.opencontainers.image.source="https://github.com/qclaogui/gaip" \
       org.opencontainers.image.description="Practices for implementing Google API Improvement Proposals (AIP) in Go."

COPY --from=builder bin/gaip_linux_amd64 /bin/gaip

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bin/gaip"]