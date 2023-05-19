package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateGitignore(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source file's path relative to the binary
	sourceGitignoreFilePath := filepath.Join(filepath.Dir(binaryPath), "app", "base", "gitignore", ".gitignore")

	// Define the destination path
	gitignoreFilePath := filepath.Join(".gitignore")

	// Use CopyFile function to copy .gitignore file
	err = utils.CopyFile(sourceGitignoreFilePath, gitignoreFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying .gitignore file: %v", err)
	}

	fmt.Printf(".gitignore file generated successfully at %s.\n", gitignoreFilePath)
	return nil
}
