package installations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ignaciojeria/einar/cmd/domain"
	"github.com/Ignaciojeria/einar/cmd/utils"
)

func InstallCommand(project string, commandName string) error {

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
	if err != nil {
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

	sourceDir := filepath.Join(templateFolderPath, installCommand.SourceDir)
	destDir := filepath.Join( /*project*/ "", installCommand.SourceDir)

	err = utils.CopyDirectory(sourceDir, destDir, []string{`"archetype`, "${project}"}, []string{`"` + project, project})
	if err != nil {
		return fmt.Errorf("error cloning %s directory: %v", commandName, err)
	}

	fmt.Printf("%s directory cloned successfully to %s.\n", commandName, destDir)

	for _, lib := range installCommand.Libraries {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = "" /*project*/
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error installing %s library %s: %v", commandName, lib, err)
		}
	}

	if err := addInstallationInsideCli( /*"project"*/ "", commandName); err != nil {
		return fmt.Errorf("failed to update .einar.template.json: %v", err)
	}

	setupFilePath := filepath.Join( /*project*/ "", "app/shared/archetype/setup.go")

	err = utils.AddImportStatement(setupFilePath, fmt.Sprintf(project+"/"+installCommand.SourceDir))
	if err != nil {
		return fmt.Errorf("failed to add import statement to setup.go: %v", err)
	}

	firstLevelDirs, err := utils.ListFirstLevelDirs(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to list first level directories: %v", err)
	}

	for _, v := range firstLevelDirs {
		err = utils.AddImportStatement(setupFilePath, fmt.Sprintf(project+"/"+installCommand.SourceDir+"/"+v))
		if err != nil {
			return fmt.Errorf("failed to add import statement to setup.go: %v", err)
		}
	}


	return nil
}
