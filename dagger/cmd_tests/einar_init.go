package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarInit(ctx context.Context, client *dagger.Client) error {
	src := client.Host().Directory("./host_output")
	container :=client.
	Container().
	From("golang:latest").
	WithDirectory("/src",src).
	WithWorkdir("/src").
	WithExec([]string{"./einar","init","my-project","https://github.com/Ignaciojeria/einar-cli-template","no-auth"})

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
