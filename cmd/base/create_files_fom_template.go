package base

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ignaciojeria/einar/cmd/domain"
	"github.com/Ignaciojeria/einar/cmd/utils"
)

func CreateFilesFromTemplate(templateFilePath string, project string) error {

	fmt.Println(templateFilePath)
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

	// Iterate over the Files slice
	for _, file := range template.BaseTemplate.Files {
		// Construct the source and destination paths
		sourcePath := filepath.Join(templateFilePath, file.SourceFile)
		destinationPath := file.DestinationFile

		// Copy the file
		err = utils.CopyFile(sourcePath, destinationPath, []string{`"archetype`, "${project}"}, []string{`"` + project, project})
		if err != nil {
			return fmt.Errorf("error copying file from %s to %s: %v for project %v", sourcePath, destinationPath, err, project)
		}

		fmt.Printf("File copied successfully from %s to %s.\n", sourcePath, destinationPath)
	}

	return nil
}
