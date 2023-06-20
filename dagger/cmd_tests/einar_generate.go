package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarGenerate(ctx context.Context, container *dagger.Container) (*dagger.Container,error) {
	// define the application einar init command
	binary := "einar"

	type GenerateCommand struct{
		Type string
		Name string
	}

	installations := []GenerateCommand{
		{
		Type : "get-controller",
		Name: "get-customer",
		},
	}

	for _, v := range installations {
		container = container.WithExec([]string{
			binary, 
			"generate",
			v.Type,
			v.Name,})
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