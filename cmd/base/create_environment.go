package base

import (
	"archetype/cmd/utils"
	"fmt"
	"path/filepath"
)

func CreateEnvironment(project string) error {
	// Define the source and destination paths
	sourceEnvFilePath := "cmd/base/environment/.environment"
	envFilePath := filepath.Join(project, ".env")

	// Use CopyFile function to copy and process .environment file
	err := utils.CopyFile(sourceEnvFilePath, envFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying .environment file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".env file generated successfully at %s.\n", envFilePath)
	return nil
}
