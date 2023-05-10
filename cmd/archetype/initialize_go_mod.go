package archetype

import (
	"fmt"
	"os/exec"
)

func InitializeGoModule(moduleName string, dependencies []string) error {
	// Initialize a new Go module
	goModCmd := exec.Command("go", "mod", "init", moduleName)
	goModCmd.Dir = moduleName
	err := goModCmd.Run()
	if err != nil {
		err := fmt.Errorf("error initializing go module for module %s: %s", moduleName, err)
		fmt.Println(err)
		return err
	}

	// Get dependencies
	for _, dependency := range dependencies {
		goGetCmd := exec.Command("go", "get", dependency)
		goGetCmd.Dir = moduleName
		err := goGetCmd.Run()
		if err != nil {
			err := fmt.Errorf("error getting dependency %s for module %s: %s", dependency, moduleName, err)
			fmt.Println(err)
			return err
		}
	}

	// Print success message
	fmt.Printf("Go module '%s' generated successfully with dependencies: %v.\n", moduleName, dependencies)
	return nil
}
