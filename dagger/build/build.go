package build

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/constants"
	"github.com/Ignaciojeria/einar/dagger/export"
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

	container = container.WithExec([]string{"go", "build", "-o", 
	constants.PATH+
	constants.BINARY_NAME, "main.go"}) // Notice the added "main.go"
	
	container = container.WithEnvVariable("BUILD_DIRECTORY",os.Getenv("BUILD_DIRECTORY"))

	if err := export.Output(ctx,container);err!=nil{
		return nil,err
	}
	
	return container,nil
}
