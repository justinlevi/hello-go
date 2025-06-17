// Dagger CI/CD pipeline for hello-go project
//
// This module implements a complete CI/CD pipeline using Dagger that replaces
// the GitHub Actions workflow. It includes formatting checks, static analysis,
// testing with coverage, multi-platform builds, and Docker image building.

package main

import (
	"context"
	"fmt"
	"strings"
	"dagger/hello-go/internal/dagger"
)

type HelloGo struct{}

// CI runs the complete CI pipeline with all checks
func (m *HelloGo) CI(ctx context.Context, source *dagger.Directory) error {
	fmt.Println("🚀 Starting Dagger CI Pipeline...")
	
	// Run all checks sequentially
	fmt.Println("\n📝 Checking code formatting...")
	if err := m.Format(ctx, source); err != nil {
		return fmt.Errorf("formatting check failed: %w", err)
	}
	
	fmt.Println("\n🔍 Running static analysis...")
	if err := m.Vet(ctx, source); err != nil {
		return fmt.Errorf("go vet failed: %w", err)
	}
	
	fmt.Println("\n🧪 Running tests with coverage...")
	coverage, err := m.Test(ctx, source)
	if err != nil {
		return fmt.Errorf("tests failed: %w", err)
	}
	fmt.Printf("✅ Tests passed! %s\n", coverage)
	
	fmt.Println("\n🔨 Building binaries...")
	if err := m.Build(ctx, source); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}
	
	fmt.Println("\n🐳 Building and testing Docker image...")
	if err := m.Docker(ctx, source); err != nil {
		return fmt.Errorf("docker build/test failed: %w", err)
	}
	
	fmt.Println("\n✨ CI Pipeline completed successfully!")
	return nil
}

// Format checks code formatting with gofmt
func (m *HelloGo) Format(ctx context.Context, source *dagger.Directory) error {
	output, err := m.baseEnv(source).
		WithExec([]string{"sh", "-c", "gofmt -l . | grep -E '\\.go$' | grep -v '^dagger/'"}).
		Stdout(ctx)
	
	if err == nil && strings.TrimSpace(output) != "" {
		return fmt.Errorf("the following files need formatting:\n%s", output)
	}
	
	// If grep returns error, it means no files need formatting (which is good)
	if err != nil && !strings.Contains(err.Error(), "exit code: 1") {
		return err
	}
	
	fmt.Println("✅ All files are properly formatted")
	return nil
}

// Vet runs go vet for static analysis
func (m *HelloGo) Vet(ctx context.Context, source *dagger.Directory) error {
	output, err := m.baseEnv(source).
		WithExec([]string{"go", "vet", "./..."}).
		Stdout(ctx)
	
	if err != nil {
		return fmt.Errorf("go vet failed: %w\n%s", err, output)
	}
	
	fmt.Println("✅ Static analysis passed")
	return nil
}

// Test runs tests with race detection and coverage
func (m *HelloGo) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	container := m.baseEnv(source).
		WithExec([]string{"go", "test", "-v", "-race", "-coverprofile=coverage.txt", "-covermode=atomic", "./..."})
	
	output, err := container.Stdout(ctx)
	if err != nil {
		return "", err
	}
	
	// Extract coverage percentage from output
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "coverage:") {
			parts := strings.Fields(line)
			for i, part := range parts {
				if part == "coverage:" && i+1 < len(parts) {
					return fmt.Sprintf("Coverage: %s", parts[i+1]), nil
				}
			}
		}
	}
	
	return "Tests passed", nil
}

// Build creates binaries for multiple platforms
func (m *HelloGo) Build(ctx context.Context, source *dagger.Directory) error {
	base := m.baseEnv(source)
	
	platforms := []struct {
		os   string
		arch string
		ext  string
	}{
		{"linux", "amd64", ""},
		{"darwin", "amd64", ""},
		{"windows", "amd64", ".exe"},
	}
	
	for _, platform := range platforms {
		binary := fmt.Sprintf("hello-api-%s-%s%s", platform.os, platform.arch, platform.ext)
		fmt.Printf("  Building %s...\n", binary)
		
		_, err := base.
			WithEnvVariable("GOOS", platform.os).
			WithEnvVariable("GOARCH", platform.arch).
			WithEnvVariable("CGO_ENABLED", "0").
			WithExec([]string{"go", "build", "-o", binary, "."}).
			Sync(ctx)
		
		if err != nil {
			return fmt.Errorf("failed to build for %s/%s: %w", platform.os, platform.arch, err)
		}
	}
	
	fmt.Println("✅ All binaries built successfully")
	return nil
}

// Docker builds and tests the Docker image
func (m *HelloGo) Docker(ctx context.Context, source *dagger.Directory) error {
	// Build the Docker image
	fmt.Println("  Building Docker image...")
	container := dag.Container().
		Build(source)
	
	// Export the image to verify it built correctly
	_, err := container.Export(ctx, "/tmp/hello-api.tar")
	if err != nil {
		return fmt.Errorf("failed to build Docker image: %w", err)
	}
	
	// Start the container as a service
	fmt.Println("  Starting container for testing...")
	service := container.
		WithExposedPort(8080).
		AsService()
	
	// Test the health endpoint
	fmt.Println("  Testing health endpoint...")
	_, err = dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "--no-cache", "curl"}).
		WithServiceBinding("hello-api", service).
		WithExec([]string{"sh", "-c", "sleep 5 && curl -f http://hello-api:8080/health"}).
		Stdout(ctx)
	
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	
	fmt.Println("✅ Docker image built and tested successfully")
	return nil
}

// baseEnv returns a container with Go environment and source code
func (m *HelloGo) baseEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.21").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod")).
		WithMountedCache("/root/.cache/go-build", dag.CacheVolume("go-build")).
		WithEnvVariable("CGO_ENABLED", "1")
}

// Shortcuts for individual operations

// CheckFormat runs only the formatting check
func (m *HelloGo) CheckFormat(ctx context.Context, source *dagger.Directory) error {
	return m.Format(ctx, source)
}

// CheckVet runs only the static analysis
func (m *HelloGo) CheckVet(ctx context.Context, source *dagger.Directory) error {
	return m.Vet(ctx, source)
}

// RunTests runs only the tests
func (m *HelloGo) RunTests(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.Test(ctx, source)
}

// BuildBinaries runs only the build step
func (m *HelloGo) BuildBinaries(ctx context.Context, source *dagger.Directory) error {
	return m.Build(ctx, source)
}

// BuildDocker runs only the Docker build and test
func (m *HelloGo) BuildDocker(ctx context.Context, source *dagger.Directory) error {
	return m.Docker(ctx, source)
}