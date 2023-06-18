package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarVersion(ctx context.Context, container *dagger.Container) error {

	// define the application einar version command
	path := "output/"
	binaryName := "einar"
	container = container.WithExec(
		[]string{
		path+binaryName,
		"version"}) // Notice the added "main.go"

	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
