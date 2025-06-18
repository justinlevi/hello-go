# Hello API

[![CI](https://github.com/justinlevi/hello-go/actions/workflows/ci.yml/badge.svg)](https://github.com/justinlevi/hello-go/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/justinlevi/hello-go)](https://goreportcard.com/report/github.com/justinlevi/hello-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)

A production-ready Go HTTP JSON API server showcasing enterprise-grade development practices, modern CI/CD with Dagger, and comprehensive testing.

## Overview

This project demonstrates a complete modern Go development workflow with:
- **HTTP JSON API** with multiple endpoints and request methods
- **Comprehensive testing** with 85%+ code coverage and benchmarks
- **Modern CI/CD** using Dagger for portable, testable pipelines
- **Enterprise patterns** including structured logging, graceful shutdown, and error handling
- **Containerization** with optimized Docker images and health checks
- **Dependency management** with automated updates via Dependabot
- **Code quality** enforcement with formatting, linting, and static analysis

## Project Structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) conventions. While our simple API currently keeps all code in `main.go`, we've prepared the directory structure for future growth:

```
hello-go/
├── .github/                    # GitHub Actions workflows and configurations
│   ├── workflows/
│   │   └── ci.yml             # Dagger CI/CD pipeline
│   ├── dependabot.yml         # Automated dependency updates
│   └── WORKFLOWS.md           # CI/CD documentation
├── cmd/                       # Command-line applications
├── dagger/                    # Dagger CI/CD pipeline code
│   ├── main.go               # Dagger pipeline implementation
│   └── go.mod                # Dagger module dependencies
├── docs/                      # Project documentation
│   ├── DAGGER_REVIEW.md      # Dagger implementation details
│   └── TESTING_GITHUB_ACTIONS.md
├── internal/                  # Private application code
├── pkg/                       # Public library code
├── tasks/                     # Task management system
│   ├── complete/             # Completed development tasks
│   └── tasks-directive.md    # Task creation standards
├── test/                      # Integration tests
├── go.mod                     # Go module definition
├── main.go                    # Application entry point with full HTTP server
├── main_test.go              # Comprehensive test suite
├── Dockerfile                 # Multi-stage Docker build
├── docker-compose.yml         # Local development environment
├── k8s-deployment.yaml        # Kubernetes deployment manifests
├── dagger.json               # Dagger configuration
├── coverage.txt              # Test coverage report
├── LICENSE                   # Apache 2.0 License
├── CLAUDE.md                 # AI assistant context and guidelines
└── README.md                 # This comprehensive guide
```

Each directory contains a README.md explaining its purpose and usage guidelines.

## Getting Started

### Prerequisites

- **Go 1.21+** - Modern Go version with latest features
- **Docker** - For containerization and local development
- **Dagger** (optional) - For running CI/CD pipelines locally

### Running Locally

1. Clone the repository
2. Install dependencies (if any):
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. Test the API:
   ```bash
   curl http://localhost:8080/hello
   ```

### Running Tests

Run the comprehensive test suite:
```bash
# Run all tests with verbose output
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Generate detailed coverage report
go test ./... -coverprofile=coverage.txt -covermode=atomic
```

### Building

```bash
# Build for current platform
go build -o hello-api

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o hello-api-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o hello-api-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o hello-api-windows-amd64.exe

# Run the built binary
./hello-api
```

### Using Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build and run manually
docker build -t hello-api .
docker run -p 8080:8080 hello-api
```

## API Endpoints

### GET /hello
Returns a JSON greeting message.

**Response:**
```json
{
  "message": "Hello, World!"
}
```

### GET /hello?name={name}
Returns a personalized greeting.

**Response:**
```json
{
  "message": "Hello, {name}!"
}
```

### POST /hello
Accepts a JSON body with a name field.

**Request:**
```json
{
  "name": "Alice"
}
```

**Response:**
```json
{
  "message": "Hello, Alice!"
}
```

### GET /health
Health check endpoint for monitoring and container orchestration.

**Response:**
```json
{
  "status": "healthy"
}
```

## Testing

This project includes comprehensive testing at multiple levels:

### Test Suite Features
- **Unit tests** for all handler functions with table-driven test cases
- **HTTP integration tests** using `httptest` for realistic request/response testing
- **Benchmark tests** for performance monitoring
- **Error scenario testing** including malformed JSON, invalid content types, and edge cases
- **Middleware testing** for logging, panic recovery, and request tracking
- **Coverage reporting** with detailed metrics

### Running Specific Test Types
```bash
# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Generate coverage report
go test -coverprofile=coverage.txt ./...

# View coverage in browser
go tool cover -html=coverage.txt
```

### Test Coverage Goals
- **85%+ code coverage** maintained across all packages
- **All public functions** have corresponding tests
- **Error paths** are explicitly tested
- **Edge cases** like empty inputs and malformed data are validated

## Development

This enterprise-grade project follows modern Go development practices:

### Code Quality Standards
- **Test-Driven Development** with 85%+ code coverage requirement
- **Comprehensive testing** including unit tests, benchmarks, and table-driven tests
- **Static analysis** with `go vet` and `gofmt` enforcement
- **Error handling** with structured error responses and proper HTTP status codes
- **Input validation** with request size limits and JSON schema validation
- **Structured logging** with request IDs, timing, and proper log levels

### Production Features
- **Graceful shutdown** with signal handling and connection draining
- **Health checks** for container and Kubernetes deployments
- **Middleware patterns** for cross-cutting concerns (logging, recovery, metrics)
- **Container optimization** with multi-stage Docker builds
- **Security hardening** with non-root users and minimal attack surface

### CI/CD with Dagger

This project uses [Dagger](https://dagger.io) for portable, testable CI/CD pipelines that run identically locally and in GitHub Actions.

#### Run Complete Pipeline Locally
```bash
# Run the full CI pipeline (requires Docker)
dagger call ci --source=.
```

#### Run Individual Pipeline Steps
```bash
# Code formatting check
dagger call format --source=.

# Static analysis
dagger call vet --source=.

# Run tests with coverage
dagger call test --source=.

# Build multi-platform binaries
dagger call build --source=.

# Build and test Docker image
dagger call docker --source=.
```

#### Benefits of Dagger
- **Portable**: Same pipeline runs locally, in CI, and anywhere Docker runs
- **Fast**: Intelligent caching speeds up subsequent runs
- **Testable**: Debug CI issues locally without pushing to remote
- **Type-safe**: Pipeline defined in Go, not YAML
- **Reproducible**: Consistent results across all environments

### GitHub Actions Integration

The project uses GitHub Actions with Dagger for:
- **Automated testing** on every push and pull request
- **Multi-platform builds** for Linux, macOS, and Windows
- **Docker image building** and testing
- **Dependency updates** via Dependabot
- **Code quality enforcement** preventing merge of failing code

### Task Management System

The project includes a comprehensive task management system in `tasks/`:
- **Completed tasks** documenting development history
- **Task templates** ensuring consistent implementation
- **Behavioral contracts** defining expected functionality
- **Integration validation** ensuring components work together

### Architecture Documentation

- **`CLAUDE.md`** - Comprehensive development guidelines and AI context
- **`docs/DAGGER_REVIEW.md`** - Detailed Dagger implementation analysis
- **`.github/WORKFLOWS.md`** - CI/CD pipeline documentation
- **Task files** - Individual feature implementation details

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.

This project serves as both a production-ready API server and an educational resource demonstrating modern Go development practices.