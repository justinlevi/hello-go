# /cmd

This directory contains the main applications for this project.

## Overview

The `/cmd` directory is for your main application entry points. Each subdirectory should match the name of the executable you want to have (e.g., `/cmd/myapp`).

## Usage

When this project grows beyond a single `main.go` file, create subdirectories here:

```
cmd/
├── api/           # Main API server application
│   └── main.go
├── worker/        # Background job processor
│   └── main.go
└── cli/           # Command-line tools
    └── main.go
```

## Guidelines

- Don't put a lot of code in the `/cmd` directory
- The code here should import and invoke code from `/internal` and `/pkg` directories
- Keep it small and simple
- Each subdirectory should produce one executable
- Common application code should be in `/internal`

## Example Structure

```go
// cmd/api/main.go
package main

import (
    "hello-api/internal/server"
    "hello-api/internal/config"
)

func main() {
    cfg := config.Load()
    srv := server.New(cfg)
    srv.Run()
}
```

## When to Use

Start using this directory when:
- You need multiple executables (API server + CLI tool)
- Your `main.go` becomes too complex
- You want clear separation between entry points

For now, our simple API lives in the root `main.go`, which is perfectly fine for small projects.