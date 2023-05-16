package base

import (
	"fmt"
	"path/filepath"
)

func CreateArchetypeSetupFile(project string) error {
	// Define the source and destination paths
	sourceSetupFilePath := "app/shared/archetype/setup.go"
	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")

	// Use CopyFile function to copy setup.go file
	err := CopyFile(sourceSetupFilePath, setupFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying setup.go file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("setup.go file generated successfully at %s.\n", setupFilePath)
	return nil
}
