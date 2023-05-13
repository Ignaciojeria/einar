package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateCiFile(project string) error {
	// Read the content of the source .gitlab-ci.yml file
	sourceCiFilePath := "cmd/base/ci/.gitlab-ci.yml"
	ciFileContentBytes, err := ioutil.ReadFile(sourceCiFilePath)
	if err != nil {
		err := fmt.Errorf("error reading source .gitlab-ci.yml file at %s: %s", sourceCiFilePath, err)
		fmt.Println(err)
		return err
	}

	ciFileContent := string(ciFileContentBytes)

	// Write the content to the destination .gitlab-ci.yml file
	ciFilePath := filepath.Join(project, ".gitlab-ci.yml")
	err = ioutil.WriteFile(ciFilePath, []byte(ciFileContent), 0644)
	if err != nil {
		err := fmt.Errorf("error writing .gitlab-ci.yml file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf(".gitlab-ci.yml file generated successfully at %s.\n", ciFilePath)
	return nil
}
