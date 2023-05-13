package base

import (
	"fmt"
	"os"
)

func CreateRootDirectory(project string) error {
	// Ensure the module name is valid
	if project == "" {
		err := fmt.Errorf("please specify a name for the new project")
		fmt.Println(err)
		return err
	}

	// Create the module directory and the main.go file
	err := os.Mkdir(project, 0755)
	if err != nil {
		err = fmt.Errorf("error creating project %s: %s", project, err)
		fmt.Println(err)
		return err
	}
	return nil
}
