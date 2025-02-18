# Docker Ext Builder

This Go project allows you to build Dockerfiles and export a predefined output path content to a TAR file.

## Motivation for this tool

I've found myself in need to build several postgres extensions and export their content for in a way that the percona postgres k8s controller could ingest it. 

I know there is a [partial docker solution](https://docs.docker.com/build/building/export/), but it does not quite fit my needs. 

Feel free to suggest any kind of changes to this project.

## Features

- Build Dockerfiles
- Export Docker image content to a TAR file

## Requirements

- Go 1.16 or higher
- Docker

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/docker-ext-builder.git
    cd docker-ext-builder
    ```

2. Build the project:

    ```bash
    go build -o docker-ext-builder
    ```


> Can be used directly by go install as well

## Usage

> Check docker-ext-builder --help

## Example

```bash
./docker-ext-builder -f ./examples/Dockerfile.test -o out.tar -p /opt/out

```

## TO DO

- [ ] Export to S3 directly
