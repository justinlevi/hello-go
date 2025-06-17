# /pkg

Library code that's safe to use by external applications.

## Overview

The `/pkg` directory contains code that can be imported by other projects. This is where you put code that you're explicitly making available for others to use.

## Usage

Structure your packages logically:

```
pkg/
├── httpclient/    # Reusable HTTP client utilities
├── logger/        # Logging utilities that others can use
├── validator/     # Input validation functions
└── response/      # Standard response formats
```

## Guidelines

- Only put code here that you want to be public
- Think carefully about the API - it's hard to change later
- Well-documented with examples
- Backward compatibility matters here
- Keep dependencies minimal

## Example Structure

```go
// pkg/response/json.go
package response

import (
    "encoding/json"
    "net/http"
)

// JSON writes a JSON response with the given status code
func JSON(w http.ResponseWriter, status int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}
```

## When to Use

Start using this directory when:
- You have utilities that could benefit other projects
- You want to share code between multiple applications
- You're building a library alongside your application
- You have truly generic, reusable code

## Important Notes

1. **Think Twice**: Not all shared code belongs in `/pkg`. Sometimes `/internal` is better.
2. **API Stability**: Once others depend on this code, changes become difficult.
3. **Alternative**: Some Go projects don't use `/pkg` at all, putting packages directly in the root.

## Not Using /pkg

It's perfectly fine to NOT use this directory. Many Go projects put their packages directly in the root:

```
myproject/
├── client/
├── server/
└── shared/
```

Use what makes sense for your project and team.