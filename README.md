# Jeff Kuang's Website

This repository holds the Go based website server code and content for Jeff Kuang.

## Purpose

*Practice makes perfect. - Benjamin Frankin*

This project started as a means to learn and practice our developer skills using the latest technologies. As new things are introduce, it is expected to see this repository to change over time.

## Architecture

- Software
  - Go
    - Nacelle
    - Gin
    - Prometheus
  - ReactJS
- Infrastructure
  - Docker

## Requirements

- [Go 1.14.x](https://golang.org/dl/)

## Compilation

As development progresses, there will be more details provided to test out and running the web application.

### Simple Execution

In its current state, this web service can be executed with the following command

``` bash
go run ./cmd/main.go
```

### Running from Executable

``` bash
# Unix based OS
go build -o bin/website ./cmd/...
./bin/website

# Windows
go build -o bin\website.exe .\cmd\...
.\bin\website.exe
```

## TODO

- Dockerize Image
  - Docker Compose Script
- CI/CD Automation?
  - Functional Tests
  - Unit Tests
  