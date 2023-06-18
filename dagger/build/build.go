package build

import (
	"context"
	"fmt"
	"dagger.io/dagger"
	"os"
)

func Binary(ctx context.Context, client *dagger.Client) (*dagger.Container,error) {
	fmt.Println("Building with Dagger")
	
	// get reference to the local project
	src := client.Host().Directory(".") 

	// get `golang` image
	container := client.Container().From("golang:latest")

	// mount root directory into `golang` image as '/src'
	container = container.WithDirectory("/src", src).WithWorkdir("/src") 

	// define the application build command
	path := "output/"
	binaryName := "einar"
	container = container.WithExec([]string{"go", "build", "-o", path+binaryName, "main.go"}) // Notice the added "main.go"
	
	container = container.WithEnvVariable("BUILD_DIRECTORY",os.Getenv("BUILD_DIRECTORY"))

	// get reference to build output directory in container
	output := container.Directory(path)

	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, path)
	if err != nil {
		return container,err
	}

	return container,nil
}
