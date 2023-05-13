package base

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type EinarCli struct {
	Version              string                 `json:"version"`
	InstallationCommands []InstallationCommands `json:"installation_commands"`
	InstallationsBase    []InstallationsBase    `json:"installations_base"`
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
	jsonFile, err := os.Open("cmd/release/latest/.einar.cli.latest.json")
	if err != nil {
		return EinarCli{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config EinarCli
	json.Unmarshal(byteValue, &config)

	return config, nil
}
