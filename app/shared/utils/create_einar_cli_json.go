package utils

import (
	"encoding/json"
	"os"

	"github.com/Ignaciojeria/einar/app/domain"
)

func CreateEinarCLIJSON(cli domain.EinarCli) error {
	cliJSON, err := json.MarshalIndent(cli, "", "  ")
	if err != nil {
		return err
	}
	fileName := ".einar.cli.json"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(cliJSON)
	return err
}
