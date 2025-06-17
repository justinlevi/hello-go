package main

import (
	"dagger/hello-go/internal/dagger"
)

// Start and return an HTTP service
func (m *HelloGo) HttpService() *dagger.Service {
	return dag.Container().
		From("python").
		WithWorkdir("/srv").
		WithNewFile("index.html", "Hello, world!").
		WithExposedPort(8080).
		AsService(dagger.ContainerAsServiceOpts{Args: []string{"python", "-m", "http.server", "8080"}})
}