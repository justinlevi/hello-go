# /internal

Private application code. This is the code you don't want others importing in their applications or libraries.

## Overview

The `/internal` directory is a special directory recognized by the Go compiler. Go will prevent any package from being imported if the import path contains an `internal` directory, unless the importing code shares a common ancestor.

## Usage

Put your private application code here:

```
internal/
├── config/        # Configuration loading and validation
├── handlers/      # HTTP request handlers
├── middleware/    # HTTP middleware
├── models/        # Internal data models
├── repository/    # Data access layer
├── service/       # Business logic
└── utils/         # Internal utilities
```

## Guidelines

- This is your app's private code
- Other projects cannot import from here
- Feel free to structure this however makes sense for your app
- Keep packages focused and cohesive
- Avoid circular dependencies

## Example Structure

```go
// internal/handlers/hello.go
package handlers

import (
    "encoding/json"
    "net/http"
    "hello-api/internal/models"
)

type HelloHandler struct {
    // dependencies
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // handler logic
}
```

## When to Use

Start using this directory when:
- Your `main.go` has too many functions
- You have code that shouldn't be imported by others
- You need clear internal structure
- You want to organize by feature or layer

## Note

You can have multiple `internal` directories at any level of your project tree. The Go compiler will enforce the import restrictions based on the directory structure.