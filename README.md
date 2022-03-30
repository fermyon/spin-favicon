# Spin-Favicon: A trivial favicon server

This TinyGo-based server serves the `/favicon.ico` request for a webserver. It is a minimal example of writing a Spin application with WASI tools.

## Building and Running

Prerequisites:

- A recent version of Go
- A recent version of TinyGo
- A recent version of [Spin](https://github.com/fermyon/spin)

Use `make build` to build, and `make serve` to start a Spin instance serving `/favicon.ico`.
Use `CTRL-C` to stop the server.

> At the time of this writing, Spin has not been released publicly, so the `go.mod` file has a `replace` line. When Spin is available publicly, the `replace` line should be removed. In the meantime, you will need to fetch `spin` from source and then adjust the path in `go.mod`.

## Configuration

By default, this looks for a `favicon.ico` file in the location `/favicon.ico` (see `spin.toml` for file mapping).
To set this to a different path, use `spin up -e FAVICON_PATH=/path/to/favicon.ico`