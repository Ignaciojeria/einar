package pipelines

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func Build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory("./einar")

	// get `golang` image
	golang := client.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/einar", src).WithWorkdir("/einar")

	// define the application build command
	path := "build/"
	golang = golang.WithExec([]string{"go", "build", "-o", path})

	// get reference to build output directory in container
	output := golang.Directory(path)

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
