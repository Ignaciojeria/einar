package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarInit(ctx context.Context, container *dagger.Container) (*dagger.Container, error) {
	// Define the einar binary name
	binary := "einar"

	// Define the project name
	projectName := "my-project"

	// Create the /output directory inside the container
	container = container.WithExec([]string{"mkdir", "-p", "/output"})

	// Set the working directory to /output
	container = container.WithWorkdir("/output")

	// Execute the einar init command
	container = container.WithExec([]string{
		binary,
		"init",
		projectName,
		"https://github.com/Ignaciojeria/einar-cli-template",
		"no-auth",
	})

	// Specify the directory in the container where einar writes its output
	containerOutputDirectory := "/output"

	// Get reference to the specified output directory in the container
	output := container.Directory(containerOutputDirectory)

	// Specify the directory on the host where you want to export the contents
	hostOutputDirectory := "host_output"

	// Export the contents of the container's output directory to the host
	_, err := output.Export(ctx, hostOutputDirectory)
	if err != nil {
		return nil, err
	}

	return container, nil
}
