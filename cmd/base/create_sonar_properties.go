package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func CreateSonarProperties(project string) error {
	envPath := filepath.Join(project, "sonar.properties")
	envContent := ""
	sonarPropertiesPath := "cmd/base/sonar/sonar.properties"
	sonarPropertiesContentBytes, err := ioutil.ReadFile(sonarPropertiesPath)
	if err != nil {
		return fmt.Errorf("error reading sonar.properties file at %s: %s", sonarPropertiesPath, err)
	}
	environmentContent := string(sonarPropertiesContentBytes)
	sonarPropertiesLines := strings.Split(environmentContent, "\n")
	for _, line := range sonarPropertiesLines {
		if strings.TrimSpace(line) != "" {
			envContent += strings.ReplaceAll(line, "${project}", project) + "\n"
		}
	}

	err = ioutil.WriteFile(envPath, []byte(envContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing sonar.properties file: %v", err)
	}
	fmt.Printf("sonar.properties file generated successfully at %s.\n", envPath)
	return nil
}
