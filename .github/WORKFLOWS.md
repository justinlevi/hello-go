# GitHub Configuration

This directory contains GitHub-specific configuration files for the project.

## Workflows

### CI Pipeline (ci.yml)

The main continuous integration pipeline that runs on every push and pull request.

**Jobs:**
1. **Test** - Runs all tests with race detection and coverage
2. **Build** - Builds the binary and tests it can start
3. **Docker** - Builds and tests the Docker image

**Features:**
- Go module caching for faster builds
- Code formatting checks (gofmt)
- Static analysis (go vet)
- Race condition detection
- Code coverage reporting
- Multi-platform builds
- Docker image testing

**Triggers:**
- Push to main branch
- All pull requests

## Dependabot Configuration

Automated dependency updates for:
- Go modules (weekly)
- Docker base images (weekly)
- GitHub Actions (weekly)

## Required Repository Settings

To ensure the CI pipeline protects your code:

1. **Branch Protection Rules** for `main`:
   - Require pull request reviews
   - Require status checks to pass:
     - `Test`
     - `Build`
     - `Docker Build`
   - Require branches to be up to date
   - Include administrators

2. **Secrets** (if using Codecov):
   - `CODECOV_TOKEN` - Get from codecov.io

## Running Actions Locally

You can test GitHub Actions locally using [act](https://github.com/nektos/act):

```bash
# Install act
brew install act

# Run all workflows
act

# Run specific job
act -j test

# Run with specific event
act pull_request
```

## Workflow Status Badge

Add to your README:
```markdown
[![CI](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/ci.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/ci.yml)
```

Replace `YOUR_USERNAME` and `YOUR_REPO` with your actual GitHub username and repository name.