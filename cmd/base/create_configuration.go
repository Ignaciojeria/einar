package base

import (
	"fmt"
	"path/filepath"
)

func CreateConfiguration(project string) error {
	// Define the source and destination directories
	sourceDir := "app/shared/config"
	destDir := filepath.Join(project, "app/shared/config")

	// Clone the source directory to the destination directory
	err := CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning config directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("config directory cloned successfully to %s.\n", destDir)
	return nil
}
