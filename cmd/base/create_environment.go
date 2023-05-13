package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func CreateEnvironment(project string) error {
	envPath := filepath.Join(project, ".env")
	envContent := ""
	environmentPath := "cmd/base/environment/.environment"
	environmentContentBytes, err := ioutil.ReadFile(environmentPath)
	if err != nil {
		return fmt.Errorf("error reading environment file at %s: %s", environmentPath, err)
	}
	environmentContent := string(environmentContentBytes)
	environmentLines := strings.Split(environmentContent, "\n")
	for _, line := range environmentLines {
		if strings.TrimSpace(line) != "" {
			envContent += strings.ReplaceAll(line, "${project}", project) + "\n"
		}
	}

	err = ioutil.WriteFile(envPath, []byte(envContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing env file: %v", err)
	}
	fmt.Printf(".env file generated successfully at %s.\n", envPath)
	return nil
}
