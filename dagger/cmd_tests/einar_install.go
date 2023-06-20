package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarInstall(ctx context.Context, container *dagger.Container) (*dagger.Container,error) {
	// Define the einar binary name
	binary := "einar"

	// Define the installations to run
	installations := []string{
		"echo-server",
		"pubsub",
		"resty",
		"firestore",
		"postgres",
		"redis",
	}

	// Iterate over the installations and execute the einar install command
	for _, v := range installations {
		container = container.WithExec([]string{
			binary,
			"install",
			v,
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
	}

	return container,nil
}