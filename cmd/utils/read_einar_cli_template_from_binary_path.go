package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ignaciojeria/einar/cmd/domain"
)

func ReadEinarTemplateFromBinaryPath(templateFolder string) (domain.EinarTemplate, error) {
	// Construct the path to the JSON file relative to the binary
	jsonFilePath := filepath.Join(templateFolder, ".einar.template.json")

	// Read the JSON file content
	jsonContentBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return domain.EinarTemplate{}, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Unmarshal the JSON content into the EinarTemplate struct
	var template domain.EinarTemplate
	err = json.Unmarshal(jsonContentBytes, &template)
	if err != nil {
		return domain.EinarTemplate{}, fmt.Errorf("error unmarshalling JSON file: %v", err)
	}
	return template, nil
}
