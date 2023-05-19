package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadEinarCliFromBinaryPath() (EinarCli, error) {
	binaryPath, err := os.Executable()
	if err != nil {
		return EinarCli{}, fmt.Errorf("error obtaining binary path: %v", err)
	}

	cliFilePath := filepath.Join(filepath.Dir(binaryPath), "app", "base", "cli", ".einar.cli.json")

	jsonFile, err := ioutil.ReadFile(cliFilePath)
	if err != nil {
		return EinarCli{}, fmt.Errorf("error reading .einar.cli.json file: %v", err)
	}

	var config EinarCli
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return EinarCli{}, fmt.Errorf("error unmarshaling .einar.cli.json file: %v", err)
	}

	return config, nil
}
