# Testing GitHub Actions Locally with Act

This guide explains how to test our GitHub Actions workflows locally using [act](https://github.com/nektos/act).

## Prerequisites

- Docker Desktop must be installed and running
- Act installed via Homebrew: `brew install act`

## Basic Usage

### Run All Workflows

```bash
# Run the default push event
act

# Run with verbose output
act -v
```

### Run Specific Events

```bash
# Simulate a pull request
act pull_request

# Simulate a push to main
act push -b main
```

### Run Specific Jobs

```bash
# Run only the test job
act -j test

# Run only the build job
act -j build

# Run only the docker job
act -j docker
```

## Working with Our CI Pipeline

### Test the Full Pipeline

```bash
# Run all jobs in the CI workflow
act -W .github/workflows/ci.yml
```

### Test Individual Jobs

```bash
# Test only the test suite
act -j test -W .github/workflows/ci.yml

# Test only the build process
act -j build -W .github/workflows/ci.yml

# Test only Docker builds
act -j docker -W .github/workflows/ci.yml
```

### Using Different Docker Images

Act uses Docker images to simulate GitHub runners. By default, it uses smaller images which might not have all tools.

```bash
# Use medium images (more tools, larger size)
act -P ubuntu-latest=catthehacker/ubuntu:act-latest

# Use large images (closest to GitHub, very large)
act -P ubuntu-latest=catthehacker/ubuntu:full-latest
```

## Common Issues and Solutions

### 1. Docker Not Running

```bash
# Error: Cannot connect to the Docker daemon
# Solution: Start Docker Desktop
```

### 2. First Run Setup

On first run, act will ask which Docker image size to use:

```
? Please choose the default image you want to use with act:
  > Medium
    Large
    Micro
```

Choose "Medium" for a good balance of features and size.

### 3. Caching Issues

Act doesn't support GitHub Actions cache by default. Our workflow handles this gracefully.

### 4. Secrets

If your workflow needs secrets:

```bash
# Create a .secrets file (don't commit this!)
echo "CODECOV_TOKEN=your-token-here" > .secrets

# Run with secrets
act -s .secrets
```

Add `.secrets` to `.gitignore`:

```bash
echo ".secrets" >> .gitignore
```

## Debugging Workflows

### Verbose Output

```bash
# See detailed execution logs
act -v

# See very detailed logs
act -vv
```

### Dry Run

```bash
# See what would run without executing
act -n
```

### List Available Jobs

```bash
# List all jobs that would run
act -l

# List jobs for specific event
act pull_request -l
```

### Interactive Shell

```bash
# Drop into a shell when action fails
act -j test --container-options "-it"
```

## Example Commands for Our Project

```bash
# Test a typical PR workflow
act pull_request -W .github/workflows/ci.yml

# Test push to main
act push -b main -W .github/workflows/ci.yml

# Quick test run (just the test job)
act -j test

# Test with different Go version
act -j test --env GO_VERSION=1.22
```

## Tips

1. **Start Small**: Test individual jobs before running the full workflow
2. **Use Verbose Mode**: Add `-v` when debugging issues
3. **Check Docker Resources**: Ensure Docker has enough memory/CPU
4. **Clean Up**: Run `docker system prune` periodically to free space

## Differences from GitHub Actions

- No access to GitHub API (PRs, issues, etc.)
- No built-in caching (actions/cache won't work)
- Smaller runner images by default
- Some GitHub-specific features unavailable
- Local file system access

Despite these limitations, act is excellent for rapid testing of workflow logic, build processes, and most CI tasks.