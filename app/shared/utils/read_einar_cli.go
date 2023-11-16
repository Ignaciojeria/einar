package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Ignaciojeria/einar/app/domain"
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
