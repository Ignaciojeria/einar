package utils

import (
	"github.com/Ignaciojeria/einar/cmd/domain"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadEinarCli() (domain.EinarCli, error) {
	jsonFile, err := os.Open(".einar.cli.json")
	if err != nil {
		return domain.EinarCli{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config domain.EinarCli
	json.Unmarshal(byteValue, &config)

	return config, nil
}
