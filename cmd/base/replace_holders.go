package base

import (
	"os"
	"strings"
)

func replacePlaceholder(filename string, placeholder string, value string) error {
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Replace the placeholder with the value
	updatedContent := strings.ReplaceAll(string(content), placeholder, value)

	// Write the updated content back to the file
	err = os.WriteFile(filename, []byte(updatedContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
