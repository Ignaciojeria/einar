package export

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func Binary(ctx context.Context, client *dagger.Client) error {
	fmt.Println("Building with Dagger")

	// get reference to the local project
	src := client.Host().Directory(".")

	// get `golang` image
	container := client.Container().From("golang:latest")

	// mount root directory into the `golang` image as '/src'
	container = container.WithDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	container = container.WithExec([]string{"go", "build", "-o", "/myapp/einar", "main.go"})

	// Specify the directory in the container where einar writes its output
	containerOutputDirectory := "/myapp"

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
