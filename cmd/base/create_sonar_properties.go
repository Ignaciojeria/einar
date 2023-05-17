package base

import (
	"archetype/cmd/utils"
	"fmt"
	"path/filepath"
)

func CreateSonarProperties(project string) error {
	// Define the source and destination paths
	sourceSonarPropertiesFilePath := "cmd/base/sonar/sonar.properties"
	sonarPropertiesFilePath := filepath.Join(project, "sonar.properties")

	// Use CopyFile function to copy and process sonar.properties file
	err := utils.CopyFile(sourceSonarPropertiesFilePath, sonarPropertiesFilePath, project)
	if err != nil {
		err := fmt.Errorf("error copying sonar.properties file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("sonar.properties file generated successfully at %s.\n", sonarPropertiesFilePath)
	return nil
}
