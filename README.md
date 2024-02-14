# playground-wire-app

A playground for Wire.

## How-to

See [./cmd/di/](./cmd/di) for Wire code.

You can see in main that we can run 2 flavors of the container:

```sh
# Run with default container
go run ./cmd

# Run with debug env - using debug container
ENV=debug go run ./cmd
```
