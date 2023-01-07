# Build Backend
FROM golang:1.19-alpine3.17 as builder

LABEL org.opencontainers.image.title="Minetest Skin Server"
LABEL org.opencontainers.image.description="Skin server for the Minetest engine"
LABEL org.opencontainers.image.authors="AFCM <afcm.contact@gmail.com>"
LABEL org.opencontainers.image.licenses="GPL-3.0"
LABEL org.opencontainers.image.source="https://github.com/AFCMS/minetest-skin-server"

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN apk add --no-cache git=2.38.2-r0 make=4.3-r1 build-base=0.5-r3
ENV CGO_ENABLED=1
RUN go build -o minetest-skin-server .

# Build Frontend
FROM node:16 as frontend-builder
RUN mkdir /build
COPY ./frontend /frontend
WORKDIR /frontend
RUN npm install --include=dev && npm run build

# Production Image
FROM alpine:3.17 as production
RUN apk update && apk add --no-cache optipng=0.7.7-r1
COPY --from=builder /build/minetest-skin-server /
RUN mkdir -p /frontend/build
COPY --from=frontend-builder /frontend/build /frontend/build

EXPOSE 8080
CMD ["./minetest-skin-server"]