package pipelines

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func Build(ctx context.Context, client *dagger.Client) error {
	fmt.Println("Building with Dagger")
	
	// get reference to the local project
	src := client.Host().Directory(".") 

	// get `golang` image
	container := client.Container().From("golang:latest")

	// mount root directory into `golang` image as '/app'
	container = container.WithDirectory("/src", src).WithWorkdir("/src") 

	// define the application build command
	path := "build"
	container = container.WithExec([]string{"go", "build", "-o", path, "main.go"}) // Notice the added "main.go"

	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
