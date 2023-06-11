package pipelines

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func Build(ctx context.Context,client *dagger.Client) error {
	fmt.Println("Building with Dagger")
	
	// get reference to the local project
	src := client.Host().Directory("./einar")

	// get `golang` image
	container := client.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	container = container.WithDirectory("/einar", src).WithWorkdir("/einar")

	// define the application build command
	path := "build/"
	container = container.WithExec([]string{"go", "build", "-o", path})

	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
