package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type EinarCli struct {
	Version              string                 `json:"version"`
	Project              string                 `json:"project"`
	InstallationsAdded   []Installation         `json:"installations_added"`
	InstallationCommands []InstallationCommands `json:"installation_commands"`
	InstallationsBase    []InstallationsBase    `json:"installations_base"`
}

type Installation struct {
	Name      string   `json:"name"`
	Libraries []string `json:"libraries"`
}

type InstallationCommands struct {
	Name      string   `json:"name"`
	Command   string   `json:"command"`
	Libraries []string `json:"libraries"`
}

type InstallationsBase struct {
	Name    string `json:"name"`
	Library string `json:"library"`
}

func ReadEinarCli() (EinarCli, error) {
	jsonFile, err := os.Open(".einar.cli.json")
	if err != nil {
		return EinarCli{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config EinarCli
	json.Unmarshal(byteValue, &config)

	return config, nil
}

func (c *EinarCli) IsInstalled(component string) bool {
	for _, installation := range c.InstallationsAdded {
		if installation.Name == component {
			return true
		}
	}
	return false
}
