package base

import (
	"github.com/Ignaciojeria/einar/cmd/domain"
	"github.com/Ignaciojeria/einar/cmd/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFilesFromTemplate(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v for project %v", err, project)
	}

	// Construct the path to the JSON file relative to the binary
	jsonFilePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", ".einar.template.json")

	// Read the JSON file content
	jsonContentBytes, err := ioutil.ReadFile(jsonFilePath)
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
		sourcePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", file.SourceFile)
		destinationPath := os.Getenv("BUILD_DIRECTORY")+file.DestinationFile

		// Copy the file
		err = utils.CopyFile(sourcePath, destinationPath, []string{`"archetype`, "${project}"}, []string{`"` + project, project})
		if err != nil {
			return fmt.Errorf("error copying file from %s to %s: %v for project %v", sourcePath, destinationPath, err, project)
		}

		fmt.Printf("File copied successfully from %s to %s.\n", sourcePath, destinationPath)
	}

	return nil
}
