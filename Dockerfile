# syntax=docker/dockerfile:1
# check=error=true

# Build Backend
FROM --platform=$BUILDPLATFORM golang:1.24-alpine3.21 AS builder-base

LABEL org.opencontainers.image.title="Luanti Skin Server"
LABEL org.opencontainers.image.description="Skin server for the Luanti engine"
LABEL org.opencontainers.image.authors="AFCMS <afcm.contact@gmail.com>"
LABEL org.opencontainers.image.licenses="GPL-3.0-or-later"
LABEL org.opencontainers.image.source="https://github.com/AFCMS/luanti-skin-server"
LABEL io.artifacthub.package.readme-url="https://raw.githubusercontent.com/AFCMS/luanti-skin-server/refs/heads/master/README.md"
LABEL io.artifacthub.package.category="integration-delivery"
LABEL io.artifacthub.package.keywords="luanti,minetest,skins,server"
LABEL io.artifacthub.package.license="GPL-3.0-or-later"
LABEL io.artifacthub.package.maintainers='[{"name":"AFCMS","email":"afcm.contact@gmail.com"}]'


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

COPY ./common ./common
COPY ./backend ./backend

# Build Frontend
FROM --platform=$BUILDPLATFORM node:22-alpine3.21 AS builder-frontend

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

FROM --platform=$BUILDPLATFORM builder-base AS builder-prod

COPY --from=builder-frontend /frontend/dist ./backend/routes/dist

# Build with cache
# https://dev.to/jacktt/20x-faster-golang-docker-builds-289n
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" \
    --mount=type=cache,id=gobuild,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o luanti-skin-server ./backend/main.go

FROM --platform=$BUILDPLATFORM builder-base AS builder-dev

# Build with cache
# https://dev.to/jacktt/20x-faster-golang-docker-builds-289n
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" \
    --mount=type=cache,id=gobuild,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o luanti-skin-server ./backend/main.go


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

EXPOSE 8080

ENTRYPOINT ["/app/luanti-skin-server"]

# Development Image
FROM base AS development

COPY --from=builder-dev /app/luanti-skin-server /app/

# Production Image
FROM base AS production

COPY --from=builder-prod /app/luanti-skin-server /app/
