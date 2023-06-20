package build

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func Binary(ctx context.Context, client *dagger.Client) (*dagger.Container, error) {
	fmt.Println("Building with Dagger")

	// get reference to the local project
	src := client.Host().Directory(".")

	// get `golang` image
	container := client.Container().From("golang:latest")

	// mount root directory into the `golang` image as '/src'
	container = container.WithDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	container = container.WithExec([]string{"go", "build", "-o", "einar", "main.go"})

	// move the binary to a directory that is in the PATH, for example /usr/local/bin
	container = container.WithExec([]string{"mv", "einar", "/usr/local/bin/einar"})

	// Now the binary should be installed and you should be able to run it from anywhere within the container
	return container, nil
}
