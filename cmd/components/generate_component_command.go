package components

import (
	"github.com/Ignaciojeria/einar/cmd/domain"
	"github.com/Ignaciojeria/einar/cmd/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GenerateComponenteCommand(project string, componentKind string, componentName string) error {

	// read einar.cli.json
	cliPath := filepath.Join(".einar.cli.json")
	cliBytes, err := ioutil.ReadFile(cliPath)
	if err != nil {
		return fmt.Errorf("failed to read .einar.cli.json: %v", err)
	}

	var cli domain.EinarCli
	err = json.Unmarshal(cliBytes, &cli)
	if err != nil {
		return fmt.Errorf("failed to unmarshal .einar.cli.json: %v", err)
	}

	templateFolderPath, err := utils.GetTemplateFolderPath(cli.Template.URL)
	if err !=nil{
		return err
	}

	jsonFilePath := filepath.Join(templateFolderPath, ".einar.template.json")
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


	for _, v := range cli.Components {
		if v.Kind == componentKind && v.Name == componentName {
			fmt.Printf("The component '%s' for '%s' already exists.\n", componentName, componentKind)
			return errors.New("component already exists")
		}
	}

	var dependencyIsPresent bool
	for _, dependency := range installCommand.DependsOn {
		if dependency == ""{
			dependencyIsPresent = true
			break
		}
		for _, installation := range cli.Installations {
			if dependency == installation.Name {
				dependencyIsPresent = true
				break
			}
		}
	}

	if !dependencyIsPresent {
		fmt.Println("Some dependencies are missing. Please install the following dependencies:")
		for _, v := range installCommand.DependsOn {
			fmt.Println("einar install " + v)
		}
		return errors.New("dependencies are not present")
	}

	setupFilePath := filepath.Join("app/shared/archetype/setup.go")

	if err := addComponentInsideCli( componentKind, componentName); err != nil {
		return fmt.Errorf("failed to update .einar.template.json: %v", err)
	}

	// Iterate over the Files slice
	for _, file := range installCommand.ComponentFiles {

		// Extract the final component name and construct the nested folder structure
		componentParts := strings.Split(componentName, "/")
		nestedFolders := strings.Join(componentParts[:len(componentParts)-1], "/")
		if nestedFolders != "" {
			nestedFolders += "/"
		}
		componentName = componentParts[len(componentParts)-1]
		destinationDirParts := strings.Split(file.DestinationDir, "/")
		baseFolder := destinationDirParts[0]
		// Remove the first folder from Dir
		file.DestinationDir =  strings.TrimPrefix(file.DestinationDir, baseFolder+"/")
		file.Port.DestinationDir = strings.TrimPrefix(file.Port.DestinationDir, baseFolder+"/")

		if file.IocDiscovery {
			err = utils.AddImportStatement(setupFilePath, fmt.Sprintf(cli.Project+"/"+baseFolder+"/"+nestedFolders+file.DestinationDir))
		}

		if file.IocDiscovery && err != nil {
			return fmt.Errorf("failed to add import statement to setup.go: %v", err)
		}

		// Construct the source and destination paths
		sourcePath := filepath.Join(templateFolderPath, file.SourceFile)
		destinationPath := baseFolder+"/"+nestedFolders+file.DestinationDir + "/" + utils.ConvertStringCase(componentName, "snake_case") + ".go"

		
		placeHolders := []string{`"archetype`}
		placeHoldersReplace := []string{`"` + project}

		if file.Port.DestinationDir != ""{
			placeHolders = []string{`"archetype`,project+"/"+baseFolder+"/"+file.Port.DestinationDir}
			placeHoldersReplace = []string{`"` + project,project+"/"+baseFolder+"/"+nestedFolders+file.Port.DestinationDir}
		}

		for _, v := range file.ReplaceHolders {
			placeHolders = append(placeHolders, v.Name)
			placeHoldersReplace = append(placeHoldersReplace,
				v.AppendAtStart+
					utils.ConvertStringCase(componentName, v.Kind)+
					v.AppendAtEnd)
		}
		// Copy the file

		if file.Port.SourceFile != "" {
			sourcePath := filepath.Join(templateFolderPath, file.Port.SourceFile)
			destinationPath := baseFolder+"/"+nestedFolders+file.Port.DestinationDir + "/" + utils.ConvertStringCase(componentName, "snake_case") + ".go"
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
