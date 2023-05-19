package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateDockerFile(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source file's path relative to the binary
	sourceDockerFilePath := filepath.Join(filepath.Dir(binaryPath), "cmd", "base", "dockerfile", "Dockerfile")

	// Define the destination path
	dockerFilePath := filepath.Join("Dockerfile")

	// Use CopyFile function to copy Dockerfile
	err = utils.CopyFile(sourceDockerFilePath, dockerFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying Dockerfile: %v", err)
	}

	fmt.Printf("Dockerfile generated successfully at %s.\n", dockerFilePath)
	return nil
}
