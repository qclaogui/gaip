# syntax=docker/dockerfile:1

##
## Deploy the application binary into a lean image
##
FROM       alpine:3.19
ARG        EXTRA_PACKAGES
RUN        apk add --no-cache ca-certificates tzdata $EXTRA_PACKAGES

LABEL org.opencontainers.image.title="gaip" \
      org.opencontainers.image.source="https://github.com/qclaogui/golang-api-server" \
      org.opencontainers.image.description="Practices for implementing Google API Improvement Proposals (AIP) in Go."

COPY   bin/gaip /bin/gaip
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bin/gaip"]