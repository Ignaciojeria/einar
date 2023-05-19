package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateVersion(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source file's path relative to the binary
	sourceVersionFilePath := filepath.Join(filepath.Dir(binaryPath), "app", "base", "version", ".version")

	// Define the destination path
	versionPath := filepath.Join(project, ".version")

	// Read the version file content
	versionContentBytes, err := ioutil.ReadFile(sourceVersionFilePath)
	if err != nil {
		return fmt.Errorf("error reading version file: %v", err)
	}

	// Write the version file to the destination
	err = ioutil.WriteFile(versionPath, versionContentBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing version file: %v", err)
	}

	fmt.Printf(".version file generated successfully at %s.\n", versionPath)
	return nil
}
