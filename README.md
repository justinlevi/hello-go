# Hello API

[![CI](https://github.com/justinlevi/hello-go/actions/workflows/ci.yml/badge.svg)](https://github.com/justinlevi/hello-go/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/justinlevi/hello-go)](https://goreportcard.com/report/github.com/justinlevi/hello-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)

A simple Go HTTP JSON API server that demonstrates Go fundamentals and best practices.

## Overview

This project implements a minimal web API that returns JSON responses, teaching:
- Go modules and project structure
- net/http standard library
- JSON marshalling
- Basic handler functions and routing
- Building and running Go binaries

## Project Structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) conventions. While our simple API currently keeps all code in `main.go`, we've prepared the directory structure for future growth:

```
hello-api/
├── go.mod          # Go module definition
├── main.go         # Main application entry point
├── main_test.go    # Unit tests
├── Dockerfile      # Container configuration
├── README.md       # This file
├── cmd/            # Main applications (prepared for future use)
├── internal/       # Private application code (prepared for future use)
├── pkg/            # Library code that can be used by external applications
├── test/           # Additional external test apps and test data
└── docs/           # Design and user documents
```

Each directory contains a README.md explaining its purpose and usage guidelines.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker (optional, for containerization)

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

```bash
go test ./... -v
```

### Building

```bash
go build -o hello-api
./hello-api
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
Health check endpoint.

**Response:**
```json
{
  "status": "healthy"
}
```

## Development

This project follows Go best practices including:
- Comprehensive unit testing (85%+ coverage target)
- Structured logging
- Error handling
- Input validation
- Container support
- Graceful shutdown

### CI/CD with Dagger

This project uses [Dagger](https://dagger.io) for portable CI/CD pipelines. Run the complete pipeline locally:

```bash
dagger call ci --source=.
```

Or run individual steps:
```bash
dagger call format --source=.   # Check code formatting
dagger call test --source=.     # Run tests with coverage
dagger call build --source=.    # Build multi-platform binaries
dagger call docker --source=.   # Build and test Docker image
```

## License

This project is for educational purposes.