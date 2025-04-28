# syntax=docker/dockerfile:1
# check=error=true

# Build Backend
FROM --platform=$BUILDPLATFORM golang:1.24-alpine3.21 AS builder

LABEL org.opencontainers.image.title="Luanti Skin Converter"
LABEL org.opencontainers.image.description="Skin converter for the Luanti engine"
LABEL org.opencontainers.image.authors="AFCMS <afcm.contact@gmail.com>"
LABEL org.opencontainers.image.licenses="GPL-3.0"
LABEL org.opencontainers.image.source="https://github.com/AFCMS/luanti-skin-server"

ARG TARGETOS
ARG TARGETARCH

ENV GOCACHE=/root/.cache/go-build

# Install build dependencies
RUN apk add --no-cache git make build-base
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" go mod download
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" go mod verify

COPY . ./

# Build with cache
# https://dev.to/jacktt/20x-faster-golang-docker-builds-289n
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" \
    --mount=type=cache,id=gobuild,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o luanti-skin-converter ./cmd/main.go

FROM ghcr.io/shssoichiro/oxipng:v9.1.4 AS oxipng

# Base for the user/tz/ca-certificates
FROM alpine:3.21 AS base-alpine

RUN adduser \
    --gecos "" \
    --system \
    --no-create-home \
    --uid "900" \
    "appuser"

# Common base
FROM scratch AS base

COPY --from=base-alpine /etc/passwd /etc/passwd

COPY --from=oxipng /usr/local/bin/oxipng /usr/local/bin/oxipng

USER appuser
WORKDIR /app

COPY --from=builder /app/luanti-skin-converter /app/luanti-skin-converter

CMD ["/app/luanti-skin-converter"]
