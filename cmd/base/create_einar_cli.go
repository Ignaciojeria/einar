package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateEinarCli(project string) error {
	// Read the content of the source .gitlab-ci.yml file
	sourceCiFilePath := "cmd/release/latest/.einar.cli.latest.json"
	ciFileContentBytes, err := ioutil.ReadFile(sourceCiFilePath)
	if err != nil {
		err := fmt.Errorf("error reading source .einar.cli.latest.json file at %s: %s", sourceCiFilePath, err)
		fmt.Println(err)
		return err
	}

	ciFileContent := string(ciFileContentBytes)

	// Write the content to the destination .gitlab-ci.yml file
	ciFilePath := filepath.Join(project, ".einar.cli.json")
	err = ioutil.WriteFile(ciFilePath, []byte(ciFileContent), 0644)
	if err != nil {
		err := fmt.Errorf("error writing .einar.cli.json file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".einar.cli.json file generated successfully at %s.\n", ciFilePath)
	return nil
}
