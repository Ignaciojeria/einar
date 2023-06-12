package pipelines

import (
	"context"
	"fmt"
	"os"
	"dagger.io/dagger"
)

func BuildRelease(ctx context.Context,client *dagger.Client) error {
	fmt.Println("Building with Dagger")
	
	// get reference to the local project
	src := client.Host().Directory(".")

	// get `goreleaser` image
	container := client.Container().From("goreleaser/goreleaser:latest")

	//set environment
	container = container.WithEnvVariable("GITHUB_TOKEN",os.Getenv("GITHUB_ACCESS_TOKEN"))

	// mount cloned repository into `goreleaser` image
	container = container.WithDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	path := "dist/"
	container = container.WithExec([]string{"release","--snapshot","--config", ".goreleaser.yml"})
	
	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}