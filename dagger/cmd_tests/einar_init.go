package cmd_tests

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/export"
	"dagger.io/dagger"
)

func EinarInit(ctx context.Context, container *dagger.Container) (*dagger.Container,error) {

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

	if err := export.Output(ctx,container);err!=nil{
		return nil,err
	}

	return container,nil
}