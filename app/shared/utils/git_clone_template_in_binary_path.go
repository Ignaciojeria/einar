package utils

import (
	"fmt"
	"os"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitCloneTemplateInBinaryPath(repositoryUrl string, userCreds string) {
	targetPath, err := GetTemplateFolderPath(repositoryUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup auth if userCreds is provided
	var auth *http.BasicAuth
	if userCreds != "no-auth" {
		user, accessToken, invalidCredsErr := SplitCredentials(userCreds)
		if invalidCredsErr != nil {
			fmt.Println("Failed to parse user credentials:", invalidCredsErr)
			return
		}

		auth = &http.BasicAuth{
			Username: user,
			Password: accessToken,
		}
	}

	// Check if the target directory already exists.
	if _, err = os.Stat(targetPath); err == nil {
		// Directory exists, perform a git pull
		repo, err := git.PlainOpen(targetPath)
		if err != nil {
			fmt.Println("Failed to open repository:", err)
			return
		}

		workTree, err := repo.Worktree()
		if err != nil {
			fmt.Println("Failed to fetch worktree:", err)
			return
		}

		pullOptions := &git.PullOptions{
			RemoteName: "origin",
			Auth:       auth,
		}

		err = workTree.Pull(pullOptions)
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println("Failed to pull updates from the remote:", err)
			return
		}

		fmt.Println("Repository updated:", targetPath)
		return
	}

	// If directory doesn't exist, perform a git clone
	cloneOptions := &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
		Auth:     auth,
	}

	_, err = git.PlainClone(targetPath, false, cloneOptions)
	if err != nil {
		fmt.Println("Failed to clone repository:", err)
		return
	}

	fmt.Println("Repository cloned to:", targetPath)
}