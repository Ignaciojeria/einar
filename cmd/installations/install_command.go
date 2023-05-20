package installations

import (
	"archetype/cmd/domain"
	"archetype/cmd/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func InstallCommand(project string, commandName string) error {
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

	var installCommand domain.InstallationCommand
	for _, command := range template.InstallationCommands {
		if command.Name == commandName {
			installCommand = command
			break
		}
	}

	if installCommand.Name == "" {
		return fmt.Errorf("%s command not found in .einar.template.json", commandName)
	}

	sourceDir := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", installCommand.SourceDir)
	destDir := filepath.Join(project, installCommand.DestinationDir)

	err = utils.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		return fmt.Errorf("error cloning %s directory: %v", commandName, err)
	}

	fmt.Printf("%s directory cloned successfully to %s.\n", commandName, destDir)

	for _, lib := range installCommand.Libraries {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error installing %s library %s: %v", commandName, lib, err)
		}
	}

	if err := AddInstallation(project, commandName); err != nil {
		return fmt.Errorf("failed to update .einar.template.json: %v", err)
	}

	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = utils.AddImportStatement(setupFilePath, fmt.Sprintf("archetype/app/shared/archetype/%s", strings.ReplaceAll(commandName, "-", "_")))
	if err != nil {
		return fmt.Errorf("failed to add import statement to setup.go: %v", err)
	}

	return nil
}
