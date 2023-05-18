package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateEnvironment(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the .environment file's path relative to the binary
	sourceEnvFilePath := filepath.Join(filepath.Dir(binaryPath), "cmd", "base", "environment", ".environment")

	// Define the destination path
	envFilePath := filepath.Join(project, ".env")

	// Use CopyFile function to copy and process .environment file
	err = utils.CopyFile(sourceEnvFilePath, envFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying .environment file: %v", err)
	}

	fmt.Printf(".env file generated successfully at %s.\n", envFilePath)
	return nil
}
