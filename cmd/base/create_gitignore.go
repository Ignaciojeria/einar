package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateGitignore(project string) error {
	gitignorePath := filepath.Join(project, ".gitignore")
	gitignoreContentBytes, err := ioutil.ReadFile("cmd/base/gitignore/.gitignore")
	if err != nil {
		return fmt.Errorf("error reading gitignore file: %v", err)
	}

	gitignoreContent := string(gitignoreContentBytes)

	err = ioutil.WriteFile(gitignorePath, []byte(gitignoreContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing gitignore file: %v", err)
	}

	fmt.Printf(".gitignore file generated successfully at %s.\n", gitignorePath)
	return nil
}
