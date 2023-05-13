package base

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyDirectory(srcDir string, dstDir string, project string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("error reading source directory: %v", err)
	}

	err = os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating destination directory: %v", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		fileInfo, err := entry.Info()
		if err != nil {
			return fmt.Errorf("error retrieving file info: %v", err)
		}

		if fileInfo.IsDir() {
			err = CopyDirectory(srcPath, dstPath, project)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcPath, dstPath, project)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
