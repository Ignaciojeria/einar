package components

import (
	"archetype/cmd/domain"
	"archetype/cmd/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateComponenteCommand(project string, componentKind string, componentName string) error {

	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get binary path: %v", err)
	}

	jsonFilePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", ".einar.template.json")
	jsonContentBytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v for project %v", err, project)
	}

	var template domain.EinarTemplate
	err = json.Unmarshal(jsonContentBytes, &template)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON file: %v for project %v", err, project)
	}

	var installCommand domain.ComponentCommands
	for _, command := range template.ComponentCommands {
		if command.Kind == componentKind {
			installCommand = command
			break
		}
	}

	if installCommand.Kind == "" {
		return fmt.Errorf("%s command not found in .einar.template.json", componentKind)
	}

	// read einar.cli.json
	cliPath := filepath.Join( /*project*/ "", ".einar.cli.json")
	cliBytes, err := ioutil.ReadFile(cliPath)
	if err != nil {
		return fmt.Errorf("failed to read .einar.cli.json: %v", err)
	}

	var cli domain.EinarCli
	err = json.Unmarshal(cliBytes, &cli)
	if err != nil {
		return fmt.Errorf("failed to unmarshal .einar.cli.json: %v", err)
	}

	setupFilePath := filepath.Join( /*project*/ "", "app/shared/archetype/setup.go")

	// Iterate over the Files slice
	for _, file := range installCommand.ComponentFiles {

		if file.IocDiscovery {
			err = utils.AddImportStatement(setupFilePath, fmt.Sprintf(cli.Project+"/"+file.DestinationDir))
		}

		if file.IocDiscovery && err != nil {
			return fmt.Errorf("failed to add import statement to setup.go: %v", err)
		}

		// Construct the source and destination paths
		sourcePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", file.SourceFile)
		destinationPath := file.DestinationDir + "/" + utils.ConvertStringCase(componentName, "snake_case") + ".go"

		placeHolders := []string{`"archetype`}
		placeHoldersReplace := []string{`"` + project}
		for _, v := range file.ReplaceHolders {
			placeHolders = append(placeHolders, v.Name)
			placeHoldersReplace = append(placeHoldersReplace,
				v.AppendAtStart+
					utils.ConvertStringCase(componentName, v.Kind)+
					v.AppendAtEnd)
		}
		// Copy the file

		if file.Port.SourceFile != "" {
			sourcePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", file.Port.SourceFile)
			destinationPath := file.Port.DestinationDir + "/" + utils.ConvertStringCase(componentName, "snake_case") + ".go"
			err = utils.CopyFile(sourcePath, destinationPath, placeHolders, placeHoldersReplace)
		}

		if err != nil {
			return fmt.Errorf("error copying file from %s to %s: %v for project %v", sourcePath, destinationPath, err, project)
		}

		err = utils.CopyFile(sourcePath, destinationPath, placeHolders, placeHoldersReplace)
		if err != nil {
			return fmt.Errorf("error copying file from %s to %s: %v for project %v", sourcePath, destinationPath, err, project)
		}
		fmt.Printf("File copied successfully from %s to %s.\n", sourcePath, destinationPath)
	}
	return nil
}
