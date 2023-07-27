package base

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ignaciojeria/einar/cmd/domain"
	"github.com/Ignaciojeria/einar/cmd/utils"
)

func CreateDirectoriesFromTemplate(templateFilePath string, project string) error {
	// Construct the path to the JSON file relative to the binary
	jsonFilePath := filepath.Join(templateFilePath, ".einar.template.json")

	// Read the JSON file content
	jsonContentBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v for project %v", err, project)
	}

	// Unmarshal the JSON content into the EinarTemplate struct
	var template domain.EinarTemplate
	err = json.Unmarshal(jsonContentBytes, &template)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON file: %v for project %v", err, project)
	}

	// Iterate over the Folders slice
	for _, folder := range template.BaseTemplate.Folders {
		// Construct the source and destination paths
		sourceDir := filepath.Join(templateFilePath, folder.SourceDir)
		destinationDir := folder.DestinationDir

		// Copy the directory
		err = utils.CopyDirectory(
			sourceDir, destinationDir,
			[]string{`"archetype`, "${project}"},
			[]string{`"` + project, project})

		if err != nil {
			return fmt.Errorf("error copying directory from %s to %s: %v for project %v", sourceDir, destinationDir, err, project)
		}

		fmt.Printf("Directory copied successfully from %s to %s.\n", sourceDir, destinationDir)
	}

	return nil
}
