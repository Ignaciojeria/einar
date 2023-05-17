package base

import (
	"archetype/cmd/utils"
	"fmt"
	"path/filepath"
)

func CreateContainer(project string) error {
	// Define the source and destination directories
	sourceDir := "app/shared/archetype/container"
	destDir := filepath.Join(project, "app/shared/archetype/container")

	// Clone the source directory to the destination directory
	err := utils.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning container directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("container directory cloned successfully to %s.\n", destDir)
	return nil
}
