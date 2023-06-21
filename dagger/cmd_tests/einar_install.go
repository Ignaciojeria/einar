package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarInstall(ctx context.Context, client *dagger.Client) error {
	// Define the installations to run
	installations := []string{
		"echo-server",
		"pubsub",
		"resty",
		"firestore",
		"postgres",
		"redis",
	}
	
	container :=client.
	Container().
	From("golang:latest").
	WithDirectory("/src",client.Host().Directory("./host_output")).
	WithWorkdir("/src")

	for _, v := range installations {
		container = container.WithExec([]string{"./einar","install",v})
	}

	// Specify the directory in the container where einar writes its output
	containerOutputDirectory := "/src"

	// Get reference to the specified output directory in the container
	output := container.Directory(containerOutputDirectory)
	
	// Specify the directory on the host where you want to export the contents
	hostOutputDirectory := "host_output"
	
	// Export the contents of the container's output directory to the host
	_, err := output.Export(ctx, hostOutputDirectory)
	if err != nil {
		return err
	}

	return nil
}