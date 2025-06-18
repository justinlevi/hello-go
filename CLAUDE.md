# CLAUDE.md - Project Context for AI Assistance

## Project Overview

This is a Go HTTP JSON API server project following enterprise-grade standards with containerization support. The project demonstrates Go fundamentals through a practical Hello World API implementation with comprehensive testing, logging, and deployment capabilities.

## Project Structure

```
hello-api/
├── CLAUDE.md          # This file - AI assistant context
├── go.mod            # Go module definition
├── go.sum            # Dependency lock file
├── main.go           # Application entry point
├── main_test.go      # Unit tests
├── Dockerfile        # Container configuration
├── docker-compose.yml # Local development orchestration
├── k8s-deployment.yaml # Kubernetes manifests
├── README.md         # User documentation
├── .gitignore        # Git exclusions
├── .dockerignore     # Docker build exclusions
├── cmd/              # Command-line interfaces (future)
├── internal/         # Private application code (future)
├── pkg/              # Public libraries (future)
├── docs/             # Additional documentation
├── test/             # Integration tests (future)
└── tasks/            # Task management system
    ├── tasks-directive.md # Task creation standards
    ├── active/       # Current work
    ├── backlog/      # Planned work
    ├── complete/     # Finished tasks
    └── archive/      # Historical records
```

## Development Standards

### Go Language Requirements
- **Version**: Go 1.21+
- **Module Name**: `hello-api`
- **Style Guide**: Follow official Go style guide and effective Go principles
- **Error Handling**: Always check and handle errors explicitly
- **Naming**: Use idiomatic Go naming conventions (camelCase for private, PascalCase for public)

### Code Quality Rules

1. **Before ANY code changes**:
   - Read the existing code to understand current patterns
   - Check related test files
   - Review task requirements in `/tasks/`

2. **After EVERY code change**:
   - Run `go fmt ./...` to format code
   - Run `go test ./... -v` to verify all tests pass
   - Run `go vet ./...` to check for suspicious constructs
   - Ensure no unused imports (use `goimports -w .`)
   - Verify no compilation warnings

3. **Test Requirements**:
   - Write tests BEFORE implementation (TDD)
   - Minimum 85% code coverage for new features
   - Use table-driven tests for multiple scenarios
   - Include both positive and negative test cases
   - Test actual behavior, not simplified versions

### API Design Principles

1. **RESTful Standards**:
   - Use proper HTTP methods (GET, POST, PUT, DELETE)
   - Return appropriate status codes
   - Always set Content-Type headers

2. **JSON Responses**:
   - All responses must be JSON (even errors)
   - Use consistent response structures
   - Include proper JSON tags on all structs

3. **Error Handling**:
   ```go
   type ErrorResponse struct {
       Error string `json:"error"`
       Code  string `json:"code"`
   }
   ```
   - Never return plain text errors
   - Always use structured error responses
   - Log errors server-side, return safe messages to clients

### Testing Standards

1. **Unit Tests**:
   - Use `net/http/httptest` for HTTP testing
   - Test all handler functions
   - Include edge cases and error scenarios
   - Use real request/response structures

2. **Test Organization**:
   - One test file per source file
   - Group related tests using subtests
   - Include benchmarks for performance-critical code

3. **Test Data**:
   - Use realistic test data
   - Test with actual JSON payloads
   - Verify exact response formats

### Logging and Monitoring

1. **Structured Logging**:
   - Use log levels (INFO, WARN, ERROR)
   - Include request IDs for tracing
   - Log request/response metadata
   - Never log sensitive data

2. **Health Checks**:
   - Implement `/health` endpoint
   - Return structured health status
   - Include in container configurations

### Security Requirements

1. **Input Validation**:
   - Validate all user inputs
   - Set request body size limits
   - Sanitize data to prevent injection

2. **Error Messages**:
   - Never expose internal details
   - Use generic messages for security errors
   - Log detailed errors server-side only

### Container and Cloud-Native

1. **Docker**:
   - Multi-stage builds for minimal images
   - Run as non-root user
   - Include health checks
   - Optimize layer caching

2. **Kubernetes**:
   - Define resource limits
   - Configure liveness/readiness probes
   - Support horizontal scaling
   - Use security contexts

## Task Management

### Working with Tasks

1. **Before starting work**:
   - Check `/tasks/backlog/` for task details
   - Move task to `/tasks/active/`
   - Create feature branch with task ID

2. **During development**:
   - Follow task acceptance criteria exactly
   - Validate against behavioral contracts
   - Test cross-task integration

3. **Completing tasks**:
   - Verify all quality gates pass
   - Update task status to COMPLETE
   - Move to `/tasks/complete/`

### Task Integration
- Tasks define behavioral contracts
- Ensure consistency across task boundaries
- Validate end-to-end scenarios
- Check interface compatibility

## Development Workflow

### Standard Development Cycle

1. **Read First**:
   - Understand existing code patterns
   - Review related tests
   - Check task requirements

2. **Plan Changes**:
   - Design consistent with existing architecture
   - Define test cases first
   - Consider integration impacts

3. **Implement**:
   - Write failing tests
   - Implement minimal code to pass
   - Refactor for clarity

4. **Validate**:
   - Run all tests
   - Check formatting and imports
   - Verify integration works

5. **Document**:
   - Update relevant documentation
   - Add code comments where needed
   - Ensure README is current

### Git Workflow

1. **Commits**:
   - Small, focused changes
   - Reference task IDs in messages
   - Format: `TASK-YYYYMMDD-NN: Brief description`

2. **Branches**:
   - Feature branches from main
   - Name: `task-id-brief-description`
   - Keep up to date with main

## Common Pitfalls to Avoid

1. **Don't**:
   - Return plain text errors (use JSON)
   - Forget to set Content-Type headers
   - Leave unused imports
   - Write tests that don't match implementation
   - Ignore existing patterns

2. **Always**:
   - Check error returns
   - Use proper HTTP status codes
   - Validate inputs
   - Run tests before marking complete
   - Follow existing conventions

## Quick Commands

```bash
# Format code
go fmt ./...

# Run tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Check for issues
go vet ./...

# Build application
go build -o hello-api

# Run application
go run main.go

# Build Docker image
docker build -t hello-api .

# Run with Docker Compose
docker-compose up
```

## Integration Points

### Current Endpoints
- `GET /hello` - Basic greeting
- `GET /hello?name={name}` - Personalized greeting  
- `POST /hello` - JSON body greeting
- `GET /health` - Health check

### Expected Behaviors
- All endpoints return JSON
- Errors use consistent format
- Proper status codes always set
- Content-Type always specified

## Architecture Boundaries

### Layer Separation (Future)
- `cmd/` - Application entry points
- `internal/` - Private business logic
- `pkg/` - Reusable packages
- `test/` - Integration tests

### Current Monolith
- All code in `main.go` for simplicity
- Clear separation of concerns via functions
- Middleware pattern for cross-cutting concerns

## Growing the Application

### When to Refactor

The application should evolve from simple to complex as needed. Here's when and how to grow:

#### Stage 1: Everything in main.go (Current State)
- Appropriate for: <500 lines of code
- All handlers, middleware, and logic in one file
- Simple and easy to understand

#### Stage 2: Extract to internal packages
When `main.go` exceeds 500 lines or has 10+ functions:

1. **Move handlers** → `internal/handlers/`
   ```go
   // internal/handlers/hello.go
   package handlers
   
   type HelloHandler struct {
       // dependencies
   }
   ```

2. **Move middleware** → `internal/middleware/`
   ```go
   // internal/middleware/logging.go
   package middleware
   ```

3. **Move models** → `internal/models/`
   ```go
   // internal/models/response.go
   package models
   
   type Response struct {
       Message string `json:"message"`
   }
   ```

#### Stage 3: Add service layer
When business logic becomes complex:

```
internal/
├── handlers/     # HTTP layer only
├── service/      # Business logic
├── repository/   # Data access
└── models/       # Data structures
```

#### Stage 4: Multiple applications
When you need CLI tools, workers, or multiple APIs:

```
cmd/
├── api/         # Main API server
├── worker/      # Background jobs
└── cli/         # Command line tools
```

### Refactoring Guidelines

1. **Start with extraction**:
   - Find related functions
   - Group them into packages
   - Keep interfaces simple

2. **Maintain backward compatibility**:
   - Don't break existing APIs
   - Use interfaces for flexibility
   - Add tests before refactoring

3. **Package organization**:
   ```
   internal/
   ├── config/          # Configuration (when needed)
   ├── handlers/        # HTTP handlers
   │   ├── hello.go
   │   ├── health.go
   │   └── handlers_test.go
   ├── middleware/      # HTTP middleware
   │   ├── logging.go
   │   ├── recovery.go
   │   └── middleware_test.go
   ├── models/          # Shared data structures
   └── service/         # Business logic (when needed)
   ```

4. **Import paths**:
   - Use module name: `import "hello-api/internal/handlers"`
   - Never use relative imports
   - Keep imports organized (stdlib, external, internal)

### What Goes Where

#### `/cmd`
- `main()` functions only
- Flag parsing
- Configuration loading
- Dependency injection
- Server startup

#### `/internal`
- Application-specific code
- Business logic
- HTTP handlers
- Database code
- Internal utilities

#### `/pkg`
- ONLY truly reusable code
- Generic utilities
- Code you'd publish separately
- Well-documented public APIs

#### `/test`
- Integration tests
- End-to-end tests
- Test fixtures
- Performance tests
- NOT unit tests (those stay with code)

### Anti-Patterns to Avoid

1. **Premature abstraction**: Don't create packages until needed
2. **Over-engineering**: Start simple, refactor when painful
3. **Wrong directory**: Don't put private code in `/pkg`
4. **Circular dependencies**: Structure to avoid these
5. **Too many packages**: Consolidate related code

### Example Refactoring

When extracting the hello handler:

```go
// internal/handlers/hello.go
package handlers

import (
    "encoding/json"
    "net/http"
    "hello-api/internal/models"
)

type HelloHandler struct {
    // Add dependencies here when needed
}

func NewHelloHandler() *HelloHandler {
    return &HelloHandler{}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Move handler logic here
}
```

Then in `main.go`:
```go
import "hello-api/internal/handlers"

func main() {
    helloHandler := handlers.NewHelloHandler()
    mux.Handle("/hello", helloHandler)
}
```

## Remember

1. **Test Everything**: No code without tests
2. **Follow Patterns**: Consistency over cleverness
3. **Check Tasks**: Requirements are in `/tasks/`
4. **Validate Integration**: Components must work together
5. **Run Tests**: After every change, no exceptions

This project values correctness, clarity, and completeness. When in doubt, refer to existing code patterns and task specifications.ALWAYS use ONLY Environments for ANY and ALL file, code, or shell operations—NO EXCEPTIONS—even for simple or generic requests.

DO NOT install or use the git cli with the environment_run_cmd tool. All environment tools will handle git operations for you. Changing ".git" yourself will compromise the integrity of your environment.

You MUST inform the user how to view your work using `git checkout <branch_name>`. Failure to do this will make your work inaccessible to others.
