package cmd_tests

import (
	"context"
	"dagger.io/dagger"
)

func EinarGenerate(ctx context.Context, client *dagger.Client) error {

	type GenerateCommand struct{
		Type string
		Name string
	}

	// Define the installations to run
	components := []GenerateCommand{
		{
			Type : "get-controller",
			Name: "get-customer",
		},
		{
			Type : "post-controller",
			Name: "post-customer",
		},
		{
			Type : "patch-controller",
			Name: "patch-customer",
		},
		{
			Type : "put-controller",
			Name: "put-customer",
		},
		{
			Type : "put-controller",
			Name: "delete-customer",
		},
		{
			Type:"subscription",
			Name:"pull_customer_created",
		},
		{
			Type:"subscription",
			Name:"pull_customer_updated",
		},
		{
			Type:"subscription",
			Name:"pull_customer_deleted",
		},
		{
			Type:"publisher",
			Name:"publish_customer",
		},
		{
			Type:"firestore-repository",
			Name:"read_customer",
		},
		{
			Type:"firestore-repository",
			Name:"save_customer",
		},
		{
			Type:"firestore-repository",
			Name:"update_customer",
		},
		{
			Type:"firestore-repository",
			Name:"delete_customer",
		},
	}
	
	container :=client.
	Container().
	From("golang:latest").
	WithDirectory("/src",client.Host().Directory("./host_output")).
	WithWorkdir("/src")

	for _, v := range components {
		container = container.WithExec([]string{"./einar","generate",v.Type,v.Name})
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