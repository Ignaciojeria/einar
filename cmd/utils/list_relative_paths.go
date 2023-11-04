package utils

import (
	"os"
	"path/filepath"
)

// ListRelativePaths scans the directory at dirPath and returns a list of all directories
// within it, relative to dirPath, with forward slashes, excluding file names.
func ListRelativePaths(dirPath string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // can't walk here,
		}
		// Ignore the root directory
		if path == dirPath {
			return nil
		}
		// Check if it's a directory; if so, get the relative path
		if d.IsDir() {
			relativePath, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err // some error occurred while getting relative path
			}
			// Convert backslashes to forward slashes for consistency across platforms
			relativePath = filepath.ToSlash(relativePath)
			paths = append(paths, relativePath)
		}
		return nil
	})

	if err != nil {
		return nil, err // filepath.WalkDir had an issue with the provided dirPath
	}

	return paths, nil
}
