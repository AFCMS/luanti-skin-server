# syntax=docker/dockerfile:1
# check=error=true

# Build Backend
FROM --platform=$BUILDPLATFORM golang:1.24-alpine3.21 AS builder

LABEL org.opencontainers.image.title="Luanti Skin Server"
LABEL org.opencontainers.image.description="Skin server for the Luanti engine"
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
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o luanti-skin-server .

# Build Frontend
FROM --platform=$BUILDPLATFORM node:22-alpine3.21 AS frontend-builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable pnpm

WORKDIR /frontend

COPY ./frontend/package.json ./
COPY ./frontend/pnpm-lock.yaml ./
COPY ./frontend/pnpm-workspace.yaml ./
RUN --mount=type=cache,id=pnpm,target="/pnpm/store" pnpm install --frozen-lockfile

COPY ./frontend ./
RUN pnpm run build

FROM ghcr.io/shssoichiro/oxipng:v9.1.5 AS oxipng

# Base for the user/tz/ca-certificates
FROM alpine:3.21 AS base-alpine

RUN apk add --no-cache tzdata ca-certificates
RUN adduser \
    --gecos "" \
    --system \
    --no-create-home \
    --uid "900" \
    "appuser"

# Common base
FROM scratch AS base

COPY --from=base-alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base-alpine /usr/share/zoneinfo /usr/share/
COPY --from=base-alpine /etc/passwd /etc/passwd

COPY --from=oxipng /usr/local/bin/oxipng /usr/local/bin/oxipng

USER appuser
WORKDIR /app

# Development Image
FROM base AS development

EXPOSE 8080

COPY --from=builder /app/index.gohtml /app/
COPY --from=builder /app/luanti-skin-server /app/

CMD ["/app/luanti-skin-server"]

# Production Image
FROM base AS production

EXPOSE 8080

COPY --from=builder /app/index.gohtml /app/
COPY --from=builder /app/luanti-skin-server /app/
COPY --from=frontend-builder /frontend/dist /app/frontend/dist

CMD ["/app/luanti-skin-server"]
