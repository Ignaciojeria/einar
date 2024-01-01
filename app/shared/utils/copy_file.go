package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(srcFile string, dstFile string, placeholders []string, values []string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer in.Close()

	// Create the destination directory if it doesn't exist yet
	dstDir := filepath.Dir(dstFile)
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		err = os.MkdirAll(dstDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", dstDir, err)
		}
	}

	out, err := os.Create(dstFile)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("error syncing destination file: %v", err)
	}

	// Replace "${project}" placeholder in the copied file
	err = replacePlaceholders(dstFile, placeholders, values)
	if err != nil {
		return fmt.Errorf("error replacing placeholder in file: %v", err)
	}

	return nil
}
