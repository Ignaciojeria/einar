package base

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"path/filepath"
)

func CreateSonarProperties(project string) error {
	// Obtain the binary's path
	binaryPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error obtaining binary path: %v", err)
	}

	// Construct the source file's path relative to the binary
	sourceSonarPropertiesFilePath := filepath.Join(filepath.Dir(binaryPath), "cmd", "base", "sonar", "sonar.properties")

	// Define the destination path
	sonarPropertiesFilePath := filepath.Join("sonar.properties")

	// Use CopyFile function to copy and process sonar.properties file
	err = utils.CopyFile(sourceSonarPropertiesFilePath, sonarPropertiesFilePath, project)
	if err != nil {
		return fmt.Errorf("error copying sonar.properties file: %v", err)
	}

	fmt.Printf("sonar.properties file generated successfully at %s.\n", sonarPropertiesFilePath)
	return nil
}
