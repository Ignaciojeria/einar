package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentFolderName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	folderName := filepath.Base(dir)
	return folderName, nil
}
