package utils

import (
	"os"
	"path/filepath"
)

func ListFirstLevelDirs(dirPath string) ([]string, error) {
	var dirs []string

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	// Convert to forward slashes if on Windows or for consistency
	for i, dir := range dirs {
		dirs[i] = filepath.ToSlash(dir)
	}

	return dirs, nil
}