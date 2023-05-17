package installations

import (
	"archetype/cmd/utils"
	"encoding/json"
	"os"
)

func AddInstallation(configPath, name, library, version string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var cli utils.EinarCli
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cli); err != nil {
		return err
	}

	// add the new installation
	cli.InstallationsAdded = append(cli.InstallationsAdded, utils.Installation{
		Name:      name,
		Libraries: []string{library + /*"@"*/ "" + version},
	})

	// write the updated config back to the file
	file, err = os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // for pretty-printing
	if err := encoder.Encode(&cli); err != nil {
		return err
	}

	return nil
}
