package cmd_tests

import (
	"context"
	"dagger.io/dagger"
	"github.com/Ignaciojeria/einar/dagger/export"
)

func EinarInstall(ctx context.Context, container *dagger.Container) (*dagger.Container,error) {

	// define the application einar init command
	path := "output/"
	binaryName := "einar"

	installations := []string{
	"echo-server",
	"pubsub",
	"resty",
	"firestore",
	"postgres",
	"redis"}

	for _, v := range installations {
		container = container.WithExec([]string{
			path+binaryName, 
			"install",
			v})
		if err := export.Output(ctx,container);err!=nil{
			return nil,err
		}	
	}
	return container,nil
}