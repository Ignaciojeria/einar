package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateCiFile(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the .gitlab-ci.yml file's path relative to the binary
	sourceCiFilePath := filepath.Join(filepath.Dir(binaryPath), "cmd", "base", "ci", ".gitlab-ci.yml")

	// Define the destination path
	ciFilePath := filepath.Join(project, ".gitlab-ci.yml")

	// Use CopyFile function to copy .gitlab-ci.yml file
	err = utils.CopyFile(sourceCiFilePath, ciFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying .gitlab-ci.yml file: %v", err)
	}

	fmt.Printf(".gitlab-ci.yml file generated successfully at %s.\n", ciFilePath)
	return nil
}
