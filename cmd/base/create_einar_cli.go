package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateEinarCli(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source file's path relative to the binary
	sourceEinarCliFilePath := filepath.Join(filepath.Dir(binaryPath), ".einar.cli.json")

	// Define the destination path
	einarCliFilePath := filepath.Join(project, ".einar.cli.json")

	// Use CopyFile function to copy .einar.cli.json file
	err = utils.CopyFile(sourceEinarCliFilePath, einarCliFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying .einar.cli.json file: %v", err)
	}

	fmt.Printf(".einar.cli.json file generated successfully at %s.\n", einarCliFilePath)
	return nil
}
