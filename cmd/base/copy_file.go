package base

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcFile string, dstFile string, project string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer in.Close()

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
	err = replacePlaceholder(dstFile, "${project}", project)
	if err != nil {
		return fmt.Errorf("error replacing placeholder in file: %v", err)
	}

	return nil
}
