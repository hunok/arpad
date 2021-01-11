# arpad

![Build status](https://github.com/hunok/arpad/workflows/merge%20to%20main/badge.svg)

Arpad [ˈaːrpaːd] is a container version scanning application. It determines weather the version of the running container is outdated or not, it compares versions to tags in the container registry.

## Status

Under development.

Project status and goals described in the [project board](https://github.com/hunok/arpad/projects/1).

## Docker container details

On each push to the `main` branch we create a container tagged as `latest` and the latest short commit git hash.

To get the latest container:
```bash
docker pull ghcr.io/hunok/arpad:latest
```

## Testing

Unit tests can be run via the default go test command
```bash
go test -cover -v
```

## Similar projects

https://github.com/replicatedhq/outdated - Kubectl plugin to find and report outdated images running in a Kubernetes cluster
