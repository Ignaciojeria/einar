package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateHelmChart(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source directory's path relative to the binary
	sourceDir := filepath.Join(filepath.Dir(binaryPath), "cmd", "base", "helmchart")

	// Define the destination directory
	destDir := filepath.Join("helmchart")

	// Clone the source directory to the destination directory
	err = utils.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		return fmt.Errorf("error cloning helmchart directory: %v", err)
	}

	fmt.Printf("helmchart directory cloned successfully to %s.\n", destDir)
	return nil
}
