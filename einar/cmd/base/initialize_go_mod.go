package base

import (
	"fmt"
	"os/exec"
)

func InitializeGoModule(dependencies []string, project string) error {
	// Initialize a new Go module
	goModCmd := exec.Command("go", "mod", "init", project)
	goModCmd.Dir = ""
	err := goModCmd.Run()
	if err != nil {
		err := fmt.Errorf("error initializing go module %s", err)
		fmt.Println(err)
		return err
	}

	// Get dependencies

	for _, dependency := range dependencies {
		goGetCmd := exec.Command("go", "get", dependency)
		goGetCmd.Dir = ""
		err := goGetCmd.Run()
		if err != nil {
			err := fmt.Errorf("error getting dependency %s %s", dependency, err)
			fmt.Println(err)
			return err
		}
	}

	// Print success message
	fmt.Printf("Go module '%s' generated successfully with dependencies:", dependencies)
	return nil
}
