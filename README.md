# Luanti Skin Server

![GitHub Workflow Status](https://img.shields.io/github/checks-status/AFCMS/luanti-skin-server/master?style=flat-square)

> [!IMPORTANT]
> This server is still in development and is not ready for production use.
> Breaking changes may occur at any time.

This server is made for serving Luanti skins to Luanti servers. It is licensed under GPLv3.

-   ✅ Easy to use and powerful **API**
-   ✅ Skins compatible with both [**VoxeLibre**](https://content.luanti.org/packages/Wuzzy/mineclone2) and [**Minetest Game**](https://content.luanti.org/packages/Luanti/minetest_game)
-   ✅ Fast and reliable, thanks to [**Docker**](https://www.docker.com), [**Golang**](https://go.dev), [**Fiber**](https://gofiber.io) and [**PostgreSQL**](https://www.postgresql.org)
-   ✅ Optimised images using [**Oxipng**](https://github.com/shssoichiro/oxipng)

## Design

The server is build with the [**Go**](https://go.dev) language on-top of the [**Fiber**](https://gofiber.io) framework.

It uses also the [**GORM**](https://gorm.io) library for interacting with the database.

The frontend is build with the [**React**](https://react.dev) library, the [**Vite**](https://vite.dev) framework
and the following libraries:

-   [**TailwindCSS**](https://tailwindcss.com) for styling
-   [**HeadlessUI**](https://headlessui.com) for dialogs, combobox, etc
-   [**Heroicons**](https://heroicons.com) for most icons
-   [**React Router**](https://reactrouter.com)
-   [**React Three Fiber**](https://github.com/pmndrs/react-three-fiber) for the 3D preview of skins

## Running Server

### Development

While it's possible to develop the server without using Docker, it's much easier so only this method is documented.

#### 1. Install Docker

Follow the official guide for your OS.

-   [Ubuntu](https://docs.docker.com/engine/install/ubuntu)
-   [Debian](https://docs.docker.com/engine/install/debian)
-   [Fedora](https://docs.docker.com/engine/install/fedora)
-   [RHEL/CentOS](https://docs.docker.com/engine/install/centos)

> [!NOTE]
> The installation links are from Docker Engine, which works only under Linux.
>
> [Docker Desktop](https://www.docker.com/products/docker-desktop) can be used on Windows, MacOS and Linux.
>
> It runs a Linux VM in the background and isn't as performant as the native version, but it's easier to install and
> use.

> [!WARNING]
> You need a [BuildKit](https://docs.docker.com/build/buildkit) enabled version of Docker to build the image.
>
> In general both the image and the included Compose files use modern features of Docker and Docker Compose.

#### 2. Install NodeJS

Install NodeJS v22 (`lts/jod`) following the [instructions](https://nodejs.org) for your system. I use [**nvm**](https://github.com/nvm-sh/nvm) under Linux.

Then enable PNPM:

```shell
corepack enable pnpm
```

#### 3. Download source code

```shell
git clone https://github.com/AFCMS/luanti-skin-server && cd luanti-skin-server
```

#### 4. Configure server

```shell
cp exemple.env .env
```

Edit the `.env` file with the config you want.

A typical development config would be:

```ini
MT_SKIN_SERVER_DATABASE_LOGGING=false

MT_SKIN_SERVER_DB_HOST=db
MT_SKIN_SERVER_DB_USER=user
MT_SKIN_SERVER_DB_PASSWORD=azerty
MT_SKIN_SERVER_DB_PORT=5432
MT_SKIN_SERVER_DB_NAME=skin_server
```

#### 5. Run services

Run backend:

```shell
COMPOSE_BAKE=true docker compose -f compose.dev.yml up --build --watch
```

Run frontend:

```shell
cd frontend && pnpm install --include=dev && pnpm run dev
```

You will now have access to the app (both frontend and API) at `http://127.0.0.1:8080`. Doing changes to the frontend
files will trigger fast refresh without needing to restart the entire app.

### Production

The supported method to run the server in production is using Docker Compose:

```yaml
---
services:
    db:
        image: "postgres:17.4-alpine"
        restart: unless-stopped
        environment:
            - POSTGRES_USER=${MT_SKIN_SERVER_DB_USER}
            - POSTGRES_PASSWORD=${MT_SKIN_SERVER_DB_PASSWORD}
            - POSTGRES_DB=${MT_SKIN_SERVER_DB_NAME}
            - DATABASE_HOST=${MT_SKIN_SERVER_DB_HOST}
        expose:
            - 5432
        volumes:
            - db:/var/lib/postgresql/data

    server:
        image: ghcr.io/afcms/luanti-skin-server:master
        environment:
            - MT_SKIN_SERVER_DB_USER=${MT_SKIN_SERVER_DB_USER}
            - MT_SKIN_SERVER_DB_PASSWORD=${MT_SKIN_SERVER_DB_PASSWORD}
            - MT_SKIN_SERVER_DB_NAME=${MT_SKIN_SERVER_DB_NAME}
            - MT_SKIN_SERVER_DB_HOST=${MT_SKIN_SERVER_DB_HOST}
            - MT_SKIN_SERVER_DB_PORT=${MT_SKIN_SERVER_DB_PORT}
            - MT_SKIN_SERVER_DATABASE_LOGGING=${MT_SKIN_SERVER_DATABASE_LOGGING}
            - MT_SKIN_SERVER_OAUTH_REDIRECT_HOST=${MT_SKIN_SERVER_OAUTH_REDIRECT_HOST}
            - MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_ID=${MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_ID}
            - MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_SECRET=${MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_SECRET}
            - MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_ID=${MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_ID}
            - MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_SECRET=${MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_SECRET}
            - MT_SKIN_SERVER_FRONTEND_DEV_MODE=false
            - MT_SKIN_SERVER_VERIFICATION_GOOGLE_SEARCH_CONSOLE=${MT_SKIN_SERVER_VERIFICATION_GOOGLE_SEARCH_CONSOLE}
        ports:
            - "8080:8080"
        depends_on:
            - db

volumes:
    db:
```

It uses the [production image](https://github.com/AFCMS/luanti-skin-server/pkgs/container/luanti-skin-server) built
by the GitHub Actions workflow, which is based on `scratch` and supports `amd64` and `arm64` architectures.

```shell
docker compose up
```

You can verify that the image have been really built by the GitHub Actions workflow and find the build log using the GitHub CLI:

```shell
gh attestation verify oci://ghcr.io/afcms/luanti-skin-server:master --repo AFCMS/luanti-skin-server
```

> [!NOTE]
> The server doesn't have TLS support, to keep it as minimal as possible. Fiber don't support HTTP/2 and HTTP/3 yet anyways.
>
> TLS should be handled by a reverse proxy like [Caddy](https://caddyserver.com) or [Traefik](https://traefik.io), which support HTTP/3 and allow easy use of Let's Encrypt, Cloudflare certificates, etc.

#### Configuration

For production the server supports some more configuration variables.

##### Google Search Console Verification

The server can use the HTML tag verification method for
the [Google Search Console](https://search.google.com/search-console) (URL prefix).

You can set the `MT_SKIN_SERVER_VERIFICATION_GOOGLE_SEARCH_CONSOLE` environment variable to Google's verification token.

You can also use the DNS record method if you want, please
checkout [Google's documentation](https://support.google.com/webmasters/answer/9008080) for more information.

##### OAuth2

The server supports OAuth2 for authentication, you can set the following environment variables to enable it.

If one of the two variables (client id, client secret) for a provider are not set, OAuth2 will be disabled for that
provider.

-   `MT_SKIN_SERVER_OAUTH_REDIRECT_HOST`: the host where the OAuth2 callback will be redirected to
-   ContentDB:
    -   `MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_ID`: the OAuth2 client ID for the ContentDB API
    -   `MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_SECRET` the OAuth2 client secret for the ContentDB API
    -   `MT_SKIN_SERVER_OAUTH_CONTENTDB_URL`: the URL of the ContentDB instance, default to `https://content.luanti.org`
    -   [Create Application](https://content.luanti.org/user/apps/)
-   GitHub:
    -   `MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_ID`: the OAuth2 client ID for the GitHub API
    -   `MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_SECRET` the OAuth2 client secret for the GitHub API
    -   [Create Application](https://github.com/settings/applications/new)

## Development Tools

I recommand using either **VSCode** or **GoLand**.

There are multiple VSCode extensions marked as recommended for the workspace.
