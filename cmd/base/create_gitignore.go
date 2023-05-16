package base

import (
	"fmt"
	"path/filepath"
)

func CreateGitignore(project string) error {
	// Define the source and destination paths
	sourceGitignoreFilePath := "cmd/base/gitignore/.gitignore"
	gitignoreFilePath := filepath.Join(project, ".gitignore")

	// Use CopyFile function to copy .gitignore file
	err := CopyFile(sourceGitignoreFilePath, gitignoreFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying .gitignore file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".gitignore file generated successfully at %s.\n", gitignoreFilePath)
	return nil
}