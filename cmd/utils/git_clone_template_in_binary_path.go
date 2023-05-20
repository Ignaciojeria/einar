package utils

import (
	"fmt"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
)

func GitCloneTemplateInBinaryPath(repositoryUrl string) {
	// Determine the path of the binary.
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to determine executable path:", err)
		return
	}
	executablePath := filepath.Dir(executable)

	// Define the target path for the git clone.
	targetPath := filepath.Join(executablePath, "./einar-cli-template")

	// Check if the target directory already exists.
	_, err = os.Stat(targetPath)
	if err == nil {
		fmt.Println("The target directory already exists:", targetPath)
		return
	}

	// Clone the repository.
	_, err = git.PlainClone(targetPath, false, &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println("Failed to clone repository:", err)
		return
	}

	fmt.Println("Repository cloned to: ", targetPath)
}
