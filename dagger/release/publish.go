package release

import (
	"context"
	"os"
	"dagger.io/dagger"
)

func Publish(ctx context.Context) error {

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	if err != nil {
		return err
	}
	
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
	container = container.WithExec([]string{"release","--config", ".goreleaser.yml"})
	
	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}