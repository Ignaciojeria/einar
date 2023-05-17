package base

import (
	"archetype/cmd/utils"
	"fmt"
	"path/filepath"
)

func CreateDockerFile(project string) error {
	// Define the source and destination paths
	sourceDockerFilePath := "cmd/base/dockerfile/Dockerfile"
	dockerFilePath := filepath.Join(project, "Dockerfile")

	// Use CopyFile function to copy Dockerfile
	err := utils.CopyFile(sourceDockerFilePath, dockerFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying Dockerfile: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Dockerfile generated successfully at %s.\n", dockerFilePath)
	return nil
}
