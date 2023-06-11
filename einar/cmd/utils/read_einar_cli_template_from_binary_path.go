package utils

import (
	"github.com/Ignaciojeria/einar-cli/einar/cmd/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadEinarTemplateFromBinaryPath() (domain.EinarTemplate, error) {
	var template domain.EinarTemplate
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return template, fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the path to the JSON file relative to the binary
	jsonFilePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", ".einar.template.json")

	// Read the JSON file content
	jsonContentBytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return template, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Unmarshal the JSON content into the EinarTemplate struct
	err = json.Unmarshal(jsonContentBytes, &template)
	if err != nil {
		return template, fmt.Errorf("error unmarshalling JSON file: %v", err)
	}

	return template, nil
}
