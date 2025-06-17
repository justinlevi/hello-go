# /test

Additional external test applications and test data.

## Overview

The `/test` directory is for additional testing that goes beyond unit tests. This includes integration tests, end-to-end tests, test data, and testing utilities.

## Important Note

**Unit tests in Go do NOT go here!** Unit tests (`*_test.go` files) should live next to the code they test. This directory is for other kinds of testing.

## Usage

Organize your additional testing needs:

```
test/
├── integration/   # Integration tests
├── e2e/          # End-to-end tests
├── fixtures/     # Test data files
├── mocks/        # Mock implementations
└── performance/  # Performance/load tests
```

## Guidelines

- Unit tests stay with the code (`main_test.go` next to `main.go`)
- Put integration tests that test multiple components here
- Store large test data files here
- Keep test utilities and helpers here
- Document how to run these tests

## Example Structure

```go
// test/integration/api_test.go
package integration

import (
    "testing"
    "net/http/httptest"
)

func TestFullAPIFlow(t *testing.T) {
    // Test complete user journey
    // This might start the full server
    // Make multiple API calls
    // Verify the entire flow works
}
```

## When to Use

Start using this directory when:
- You need integration or e2e tests
- You have large test data files
- You need test utilities shared across packages
- You want to separate different types of tests

## Running Tests

```bash
# Run unit tests (not in /test)
go test ./...

# Run integration tests
go test ./test/integration/...

# Run all tests including integration
go test ./... ./test/...
```

## Test Data

If you have test data files:

```
test/
└── fixtures/
    ├── valid_request.json
    ├── invalid_request.json
    └── large_dataset.csv
```

Access them in tests:
```go
data, err := os.ReadFile("../../test/fixtures/valid_request.json")
```