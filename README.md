# IMGIX

This repository contains a Go project template configured with an advanced development environment utilizing Docker, Docker Compose, and a suite of Go development tools aimed at boosting productivity and ensuring code quality.

## Prerequisites

- [Go v1.22.3](https://go.dev)
- [Docker v26.1.1](https://www.docker.com)
- [Docker Compose v2.27.0](https://docs.docker.com/reference/cli/docker/compose)
- [Taskfile v3.37.2](https://taskfile.dev)
- Ensure all dependencies for [templ](https://templ.guide), [sqlc](https://sqlc.dev), [golangci-lint](https://golangci-lint.run), [depth](https://github.com/KyleBanks/depth), and [air](https://github.com/air-verse/air) are installed.

## Getting Started

To get started with this project, clone the repository to your local machine:

```bash
git clone https://yourrepository.com/yourproject.git
cd web-automator
```

Build the project:

```bash
task build
```

Start the Docker services in development mode:

```bash
task docker:start:dev
```

## Tasks

- `build`: Generate code from templates, generate SQL code, and compile the Go application.
- `test`: Run unit tests for the Go application.
- `format`: Format the Go source code to match the official Go formatting standards.
- `lint`: Run linters on the Go source code to identify potential issues.
- `depth`: Analyze the dependency tree of the Go application.
- `start`: Start the compiled Go application from the generated binary.
- `start:dev`: Start the Go application in development mode with live reloading using Air.
- `docker:start:prod` or `docker:start:dev`: Start Docker services using Docker Compose files based on the specified environment mode.
- `docker:stop:prod` or `docker:stop:dev`: Stop Docker services and remove any orphan containers using the specified Docker Compose files.
- `docker:remove:prod` or `docker:remove:dev`: Fully clean up Docker environment by stopping services, removing orphans, volumes, and all associated images.

## License

This project is licensed under the [MIT License](./LICENSE).
