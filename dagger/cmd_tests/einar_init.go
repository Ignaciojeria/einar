package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarInit(ctx context.Context, container *dagger.Container) error {

	// define the application einar init command
	path := "output/"
	binaryName := "einar"
	projectName := "my-project"
	container = container.WithExec([]string{
		path+binaryName, 
		"init",
		projectName,
		"https://github.com/Ignaciojeria/einar-cli-template",
		"no-auth"}) // Notice the added "main.go"

	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}