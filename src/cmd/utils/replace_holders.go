package utils

import (
	"errors"
	"os"
	"strings"
)

func replacePlaceholders(filename string, placeholders []string, values []string) error {
	if len(placeholders) != len(values) {
		return errors.New("placeholders and values arrays must have the same length")
	}
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert the file content to a string
	updatedContent := string(content)

	// Replace each placeholder with the corresponding value
	for i, placeholder := range placeholders {
		updatedContent = strings.ReplaceAll(updatedContent, placeholder, values[i])
	}

	// Write the updated content back to the file
	err = os.WriteFile(filename, []byte(updatedContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

