# syntax=docker/dockerfile:1

FROM golang:1.21.5-bullseye AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/gaip cmd/main.go

##
## Deploy the application binary into a lean image
##
FROM gcr.io/distroless/static

LABEL  org.opencontainers.image.title="gaip" \
       org.opencontainers.image.source="https://github.com/qclaogui/golang-api-server" \
       org.opencontainers.image.description="Practices for implementing Google API Improvement Proposals (AIP) in Go."

COPY --from=builder /bin/gaip /bin/gaip

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bin/gaip"]