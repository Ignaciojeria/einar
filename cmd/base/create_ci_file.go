package base

import (
	"archetype/cmd/utils"
	"fmt"
	"path/filepath"
)

func CreateCiFile(project string) error {
	// Define the source and destination paths
	sourceCiFilePath := "cmd/base/ci/.gitlab-ci.yml"
	ciFilePath := filepath.Join(project, ".gitlab-ci.yml")

	// Use CopyFile function to copy .gitlab-ci.yml file
	err := utils.CopyFile(sourceCiFilePath, ciFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying .gitlab-ci.yml file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".gitlab-ci.yml file generated successfully at %s.\n", ciFilePath)
	return nil
}
