package base

import (
	"fmt"
	"path/filepath"
)

func CreateHelmChart(project string) error {
	// Define the source and destination directories
	sourceDir := "cmd/base/helmchart"
	destDir := filepath.Join(project, "helmchart")

	// Clone the source directory to the destination directory
	err := CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning helmchart directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("helmchart directory cloned successfully to %s.\n", destDir)
	return nil
}
