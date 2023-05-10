package archetype

import (
	"fmt"
	"os"
)

func CreateRootDirectory(moduleName string) error {
	// Ensure the module name is valid
	if moduleName == "" {
		err := fmt.Errorf("please specify a name for the new module")
		fmt.Println(err)
		return err
	}

	// Create the module directory and the main.go file
	err := os.Mkdir(moduleName, 0755)
	if err != nil {
		err = fmt.Errorf("error creating directory for module %s: %s", moduleName, err)
		fmt.Println(err)
		return err
	}
	return nil
}
