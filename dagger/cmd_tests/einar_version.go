package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarVersion(ctx context.Context, container *dagger.Container) error {

	// define the application einar version command
	binary := "einar" // Use the full path to the binary
	container = container.WithExec(
		[]string{
		binary,
		"version"})

	// get reference to build output directory in container
	output := container.Directory("/output") // Use the correct path to the directory

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, "host_output") // Replace "path_on_host" with the actual path on the host
	if err != nil {
		return err
	}

	return nil
}
