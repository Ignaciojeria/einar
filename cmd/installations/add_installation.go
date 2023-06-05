package installations

import (
	"archetype/cmd/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func AddInstallation(project, commandName string) error {
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get binary path: %v", err)
	}

	// read einar.template.json
	templatePath := filepath.Join(filepath.Dir(binaryPath), "einar-cli-template", ".einar.template.json")
	templateBytes, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read .einar.template.json: %v", err)
	}

	var template domain.EinarTemplate
	err = json.Unmarshal(templateBytes, &template)
	if err != nil {
		return fmt.Errorf("failed to unmarshal .einar.template.json: %v", err)
	}

	// find the command
	var command domain.InstallationCommand
	for _, cmd := range template.InstallationCommands {
		if cmd.Name == commandName {
			command = cmd
			break
		}
	}

	if command.Name == "" {
		return fmt.Errorf("command %s not found in .einar.template.json", commandName)
	}

	// read einar.cli.json
	cliPath := filepath.Join(project, ".einar.cli.json")
	cliBytes, err := ioutil.ReadFile(cliPath)
	if err != nil {
		return fmt.Errorf("failed to read .einar.cli.json: %v", err)
	}

	var cli domain.EinarCli
	err = json.Unmarshal(cliBytes, &cli)
	if err != nil {
		return fmt.Errorf("failed to unmarshal .einar.cli.json: %v", err)
	}

	// add the command to the CLI
	cli.Installations = append(cli.Installations, domain.Installation{
		Name:      command.Name,
		Libraries: command.Libraries,
	})

	// write back the updated einar.cli.json
	cliBytes, err = json.MarshalIndent(cli, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal .einar.cli.json: %v", err)
	}

	err = ioutil.WriteFile(cliPath, cliBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write .einar.cli.json: %v", err)
	}

	return nil
}
