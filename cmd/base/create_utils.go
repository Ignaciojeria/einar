package base

import (
	"fmt"
	"path/filepath"
)

func CreateUtils(project string) error {
	// Define the source and destination directories
	sourceDir := "app/shared/utils"
	destDir := filepath.Join(project, "app/shared/utils")

	// Clone the source directory to the destination directory
	err := CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning utils directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("utils directory cloned successfully to %s.\n", destDir)
	return nil
}
