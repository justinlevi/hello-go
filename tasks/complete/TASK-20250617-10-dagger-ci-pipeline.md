# TASK-20250617-10 - Implement Dagger CI Pipeline
Status: ACTIVE
Priority: HIGH
Assignee: TBD
Created: 2025-06-17
Updated: 2025-06-17

## Progress Notes
- Initial basic implementation completed (test only)
- Need to add: formatting, vet, coverage, builds, Docker

## User Story
As a developer
I want to use Dagger for CI/CD instead of GitHub Actions
So that I have portable, testable, and programmable CI pipelines that run anywhere

## Acceptance Criteria
- Dagger pipeline implemented in Go using the Dagger Go SDK
- Pipeline performs all tasks currently in GitHub Actions:
  - Install Go dependencies
  - Run gofmt checks
  - Run go vet static analysis
  - Run tests with coverage
  - Build binary for multiple platforms
  - Build and test Docker image
- Pipeline can be run locally with `dagger run go run ./dagger`
- GitHub Actions workflow updated to use Dagger
- Pipeline is faster than pure GitHub Actions due to caching
- All existing quality gates maintained

## Gherkin Scenarios
```gherkin
Feature: Dagger CI Pipeline

Scenario: Local pipeline execution
  Given Dagger is installed locally
  When I run "dagger run go run ./dagger"
  Then all CI checks execute successfully
  And results are displayed in the terminal
  And caching speeds up subsequent runs

Scenario: GitHub Actions integration
  Given Dagger pipeline is implemented
  When code is pushed to GitHub
  Then GitHub Actions runs the Dagger pipeline
  And all checks pass as before
  And execution time is improved

Scenario: Cross-platform builds
  Given the Dagger pipeline is running
  When it reaches the build stage
  Then binaries are created for linux/amd64, darwin/amd64, and windows/amd64
  And each binary is properly named
```

## Testing & Validation Steps
- Run pipeline locally multiple times to verify caching
- Ensure all checks from GitHub Actions are replicated
- Verify pipeline works in GitHub Actions environment
- Compare execution times between old and new pipelines
- Test failure scenarios (formatting errors, test failures)
- Verify Docker build and test functionality
- Ensure coverage reports are still generated
- Test on different operating systems

## Behavioral Contract
- **Inputs**:
  - Source code from current directory
  - Go module files (go.mod, go.sum)
  - Dockerfile for container builds
- **Outputs**:
  - Pass/fail status for each check
  - Test coverage percentage
  - Built binaries for multiple platforms
  - Docker image (locally)
  - Detailed logs for debugging
- **Side Effects**:
  - Dagger cache created/updated
  - Local binaries produced (optional)
  - Coverage files generated
- **Integration Points**:
  - GitHub Actions workflow
  - Local development environment
  - Docker daemon

## Parent/Sub Tasks
- Parent Task: None
- Sub Tasks: 
  - Future: Add deployment stages
  - Future: Add release automation
  - Future: Add security scanning
- Dependencies: 
  - Dagger already initialized (dagger.json exists)
  - Docker must be running
  - Go SDK files already generated

## Cross-Task Validation
- Must run all tests from previous tasks
- Should respect .gitignore patterns
- Must handle all file types correctly
- Should integrate with existing Docker setup
- Must maintain compatibility with current tooling

## Implementation Notes

### Dagger Pipeline Structure
```go
// dagger/main.go
package main

import (
    "context"
    "dagger.io/dagger"
)

func main() {
    ctx := context.Background()
    
    // Initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
    if err != nil {
        panic(err)
    }
    defer client.Close()
    
    // Define pipeline stages:
    // 1. Source: Load project files
    // 2. Dependencies: Install Go modules
    // 3. Format: Check code formatting
    // 4. Vet: Run static analysis
    // 5. Test: Run tests with coverage
    // 6. Build: Create binaries
    // 7. Docker: Build and test container
}
```

### Key Dagger Concepts to Implement
- **Containers**: Use Go containers for builds
- **Caching**: Cache Go modules and build artifacts
- **Parallelization**: Run independent tasks concurrently
- **Multi-platform**: Build for different OS/arch combinations
- **Secrets**: Handle any required secrets properly

### GitHub Actions Integration
Update `.github/workflows/ci.yml` to use Dagger:
```yaml
- name: Run Dagger Pipeline
  uses: dagger/dagger-for-github@v6
  with:
    verb: run
    args: go run ./dagger
    cloud-token: ${{ secrets.DAGGER_CLOUD_TOKEN }} # Optional
```

### Expected Benefits
- Portable CI that runs the same locally and in CI
- Better caching and performance
- Easier debugging of CI issues
- Type-safe pipeline definition in Go
- Reusable pipeline components