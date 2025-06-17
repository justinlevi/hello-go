# Hello API

A simple Go HTTP JSON API server that demonstrates Go fundamentals and best practices.

## Overview

This project implements a minimal web API that returns JSON responses, teaching:
- Go modules and project structure
- net/http standard library
- JSON marshalling
- Basic handler functions and routing
- Building and running Go binaries

## Project Structure

```
hello-api/
├── go.mod          # Go module definition
├── main.go         # Main application entry point
├── main_test.go    # Unit tests
├── Dockerfile      # Container configuration
└── README.md       # This file
```

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

## License

This project is for educational purposes.