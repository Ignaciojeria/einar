package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateArchetypeSetupFile(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the setup.go file's path relative to the binary
	sourceSetupFilePath := filepath.Join(filepath.Dir(binaryPath), "app", "shared", "archetype", "setup.go")

	// Define the destination path
	setupFilePath := filepath.Join("app", "shared", "archetype", "setup.go")

	// Use CopyFile function to copy setup.go file
	err = utils.CopyFile(sourceSetupFilePath, setupFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying setup.go file: %v", err)
	}

	fmt.Printf("setup.go file generated successfully at %s.\n", setupFilePath)
	return nil
}
