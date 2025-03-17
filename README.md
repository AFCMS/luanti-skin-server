# Luanti Skin Server

![GitHub Workflow Status](https://img.shields.io/github/checks-status/AFCMS/luanti-skin-server/master?style=flat-square)

> [!IMPORTANT]
> This server is still in development and is not ready for production use.
> Breaking changes may occur at any time.

This server is made for serving Luanti skins to Luanti servers. It is licensed under GPLv3.

- ✅ Easy to use and powerful **API**
- ✅ Skins compatible with both **VoxeLibre** and **Minetest Game**
- ✅ Fast and reliable, thanks to **Docker**, **Golang**, **Fiber** and **PostgreSQL**
- ✅ Optimised images using **Oxipng**

## Design

The server is build with the [**Go**](https://go.dev) language on-top of the [**Fiber**](https://gofiber.io) framework.

It uses also the [**GORM**](https://gorm.io) library for interacting with the database.

The frontend is build with the [**React**](https://reactjs.org) library, the [**Vite**](https://vitejs.dev) framework
and the following libraries:

- [**TailwindCSS**](https://tailwindcss.com) for styling
- [**HeadlessUI**](https://headlessui.com) for dialogs, combobox, etc
- [**Heroicons**](https://heroicons.com) for most icons
- [**React Router**](https://reactrouter.com)
- [**React Three Fiber**](https://github.com/pmndrs/react-three-fiber) for the 3D preview of skins

## Running Server

### Development (Docker)

#### 1. Install `docker`

Follow the official guide for your OS.

- [Ubuntu](https://docs.docker.com/engine/install/ubuntu)
- [Debian](https://docs.docker.com/engine/install/debian)
- [Fedora](https://docs.docker.com/engine/install/fedora)
- [RHEL/CentOS](https://docs.docker.com/engine/install/centos)

> [!NOTE]
> The installation links are from Docker Engine, which works only under Linux.
>
> [Docker Desktop](https://www.docker.com/products/docker-desktop) can be used on Windows, MacOS and Linux.
>
> It runs a Linux VM in the background and isn't as performant as the native version, but it's easier to install and
> use.

> [!WARNING]
> You need a [BuildKit](https://docs.docker.com/build/buildkit) enabled version of Docker to build the image.

#### 2. Download source code

```shell
git clone https://github.com/AFCMS/luanti-skin-server && cd luanti-skin-server
```

#### 3. Configure server

```shell
cp exemple.env .env
```

Edit the `.env` file with the config you want.

A typical production config would be:

```ini
MT_SKIN_SERVER_DATABASE_LOGGING=false

MT_SKIN_SERVER_DB_HOST=db
MT_SKIN_SERVER_DB_USER=user
MT_SKIN_SERVER_DB_PASSWORD=azerty
MT_SKIN_SERVER_DB_PORT=5432
MT_SKIN_SERVER_DB_NAME=skin_server
```

#### 4. Run services

```shell
docker compose -f compose.dev.yml up --build
```

You will now have access to the app (both frontend and API) at `http://localhost:8080`. Doing changes to the frontend
files will trigger fast refresh without needing to restart the entire app.

### Development (host, not recommended)

It's possible to run the server without Docker, but it's not recommended.

#### 1. Install Go, NodeJS and Oxipng

Follow the official guides for you OS.

I recommend using NodeJS v20 installed using [**nvm**](https://github.com/nvm-sh/nvm) under linux.

Oxipng instructions can be found [here](https://github.com/shssoichiro/oxipng)

#### 2. Download source code

```shell
git clone https://github.com/AFCMS/luanti-skin-server && cd luanti-skin-server
```

#### 3. Install Go dependencies

```shell
go mod download
```

#### 4. Install NodeJS dependencies

```shell
cd frontend && npm install --include=dev && cd ..
```

#### 5. Configure server

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

You need a PostgreSQL database running on the given host and port.

#### 6. Frontend

The frontend served by the Fiber backend can be build before running the app and served statically, the Vite development
server can also be proxied by the backend to avoid rebuilding everytime.

Static files can be built like this before launching the server:

```shell
cd frontend && npm run build && cd ..
```

If you want to use Vite's development server, you can run it like this:

```shell
cd frontend && npm run dev
```

With the following additional configuration in the `.env` file:

```ini
MT_SKIN_SERVER_FRONTEND_DEV_MODE=true
MT_SKIN_SERVER_FRONTEND_URL=http://localhost:5173
```

> [!CAUTION]
> The Vite server configuation makes it exposed on your local network by default to make it accessible in Docker, you
> can change this
> behaviour in the Vite configuration.

#### 7. Build and run backend

```shell
go build && ./luanti-skin-server
```

### Production

There is an [exemple](https://github.com/AFCMS/luanti-skin-server/blob/master/compose.prod.yml) production Docker
Compose file in the repository.

It uses the [production image](https://github.com/AFCMS/luanti-skin-server/pkgs/container/luanti-skin-server) built
by the GitHub Actions workflow, which supports `amd64` and `arm64` architectures.

```shell
docker pull ghcr.io/afcms/luanti-skin-server:master
docker compose -f compose.prod.yml up
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

- `MT_SKIN_SERVER_OAUTH_REDIRECT_HOST`: the host where the OAuth2 callback will be redirected to
- ContentDB:
    - `MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_ID`: the OAuth2 client ID for the ContentDB API
    - `MT_SKIN_SERVER_OAUTH_CONTENTDB_CLIENT_SECRET` the OAuth2 client secret for the ContentDB API
    - `MT_SKIN_SERVER_OAUTH_CONTENTDB_URL`: the URL of the ContentDB instance, default to `https://content.luanti.org`
    - [Create Application](https://content.luanti.org/user/apps/)
- GitHub:
    - `MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_ID`: the OAuth2 client ID for the GitHub API
    - `MT_SKIN_SERVER_OAUTH_GITHUB_CLIENT_SECRET` the OAuth2 client secret for the GitHub API
    - [Create Application](https://github.com/settings/applications/new)

## Development Tools

I recommand using either **VSCode** or **GoLand**.

There are multiple VSCode extensions marked as recommended for the workspace.
