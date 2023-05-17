package base

import (
	"fmt"
	"path/filepath"
)

func CreateEinarCli(project string) error {
	// Define the source and destination paths
	sourceEinarCliFilePath := ".einar.cli.json"
	einarCliFilePath := filepath.Join(project, ".einar.cli.json")

	// Use CopyFile function to copy .einar.cli.latest.json file
	err := CopyFile(sourceEinarCliFilePath, einarCliFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying .einar.cli.latest.json file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".einar.cli.json file generated successfully at %s.\n", einarCliFilePath)
	return nil
}
