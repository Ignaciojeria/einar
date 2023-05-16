package base

import (
	"fmt"
	"path/filepath"
)

func CreateEnvironment(project string) error {
	// Define the source and destination paths
	sourceEnvFilePath := "cmd/base/environment/.environment"
	envFilePath := filepath.Join(project, ".env")

	// Use CopyFile function to copy and process .environment file
	err := CopyFile(sourceEnvFilePath, envFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying .environment file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".env file generated successfully at %s.\n", envFilePath)
	return nil
}
