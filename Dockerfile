FROM golang:1.20.6-bullseye as builder
LABEL maintainer="qclaogui@gmail.com"

WORKDIR /root

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main cmd/main.go

FROM gcr.io/distroless/static

COPY --from=builder /root/main /usr/local/bin/main

ENTRYPOINT ["main"]