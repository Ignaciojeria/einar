package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateDockerFile(project string) error {
	// Read the content of the source .gitlab-ci.yml file
	sourceCiFilePath := "cmd/base/dockerfile/Dockerfile"
	ciFileContentBytes, err := ioutil.ReadFile(sourceCiFilePath)
	if err != nil {
		err := fmt.Errorf("error reading source Dockerfile file at %s: %s", sourceCiFilePath, err)
		fmt.Println(err)
		return err
	}

	ciFileContent := string(ciFileContentBytes)

	// Write the content to the destination .gitlab-ci.yml file
	ciFilePath := filepath.Join(project, "Dockerfile")
	err = ioutil.WriteFile(ciFilePath, []byte(ciFileContent), 0644)
	if err != nil {
		err := fmt.Errorf("error writing Dockerfile file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Dockerfile file generated successfully at %s.\n", ciFilePath)
	return nil
}
