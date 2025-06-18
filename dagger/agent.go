package main

import (
	"dagger/hello-go/internal/dagger"
)

func (m *HelloGo) GoProgram(
	// The programming assignment, e.g. "write me a curl clone"
	assignment string,
) *dagger.Container {
	environment := dag.Env().
		WithStringInput("assignment", assignment, "the assignment to complete").
		WithContainerInput("builder", dag.Container().From("golang:latest").WithWorkdir("/app"), "a container to use for building Go code").
		WithContainerOutput("completed", "the completed assignment in the Golang container")

	work := dag.LLM().
	  WithEnv(environment).
	  WithPrompt(`
		You are an expert Go programmer with an assigment to create a Go program 
		Create files in the default directory in $builder 
		Always build the code to make sure it is valid 
		Do not stop until your assignment is completed and the code builds 
		Your assignment is: $assignment`)

	return work.Env().Output("completed").AsContainer()
}