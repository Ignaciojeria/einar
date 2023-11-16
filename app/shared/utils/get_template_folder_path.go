package utils

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetTemplateFolderPath(repositoryUrl string) (string, error) {
	// Determine the path of the binary.
	executable, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("Failed to determine executable path: %w", err)
	}

	executablePath := filepath.Dir(executable)

	// Parse repository URL
	u, err := url.Parse(repositoryUrl)
	if err != nil {
		return "", fmt.Errorf("Failed to parse URL: %w", err)
	}

	// Remove 'www.' if present in the host
	u.Host = strings.TrimPrefix(u.Host, "www.")

	// Create a suitable file system path from the repository URL
	repositoryPath := strings.TrimPrefix(u.Path, "/")
	repositoryPath = strings.TrimSuffix(repositoryPath, ".git")

	// Combine host and path
	repositoryPath = filepath.Join(u.Host, repositoryPath)

	// Define the target path for the git clone.
	targetPath := filepath.Join(executablePath, repositoryPath)

	return targetPath, nil
}
