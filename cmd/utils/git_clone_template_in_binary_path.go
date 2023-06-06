package utils

import (
	"fmt"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitCloneTemplateInBinaryPath(repositoryUrl string, userCreds string) {
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

	cloneOptions := &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
	}

	if userCreds != "no-auth" {
		user, accessToken, invalidCredsErr := SplitCredentials(userCreds)
		if invalidCredsErr != nil {
			fmt.Println("Failed to parse user credentials:", invalidCredsErr)
			return
		}

		cloneOptions.Auth = &http.BasicAuth{
			Username: user,
			Password: accessToken,
		}
	}

	// Clone the repository.
	_, err = git.PlainClone(targetPath, false, cloneOptions)

	if err != nil {
		fmt.Println("Failed to clone repository:", err)
		return
	}

	fmt.Println("Repository cloned to: ", targetPath)
}
