package base

import (
	"fmt"
	"os/exec"
)

func InitializeGoModule(project string, dependencies []string) error {
	// Initialize a new Go module
	goModCmd := exec.Command("go", "mod", "init", archetype)
	goModCmd.Dir = project
	err := goModCmd.Run()
	if err != nil {
		err := fmt.Errorf("error initializing go module for project %s: %s", project, err)
		fmt.Println(err)
		return err
	}

	// Get dependencies
	for _, dependency := range dependencies {
		goGetCmd := exec.Command("go", "get", dependency)
		goGetCmd.Dir = project
		err := goGetCmd.Run()
		if err != nil {
			err := fmt.Errorf("error getting dependency %s for project %s: %s", dependency, project, err)
			fmt.Println(err)
			return err
		}
	}

	// Print success message
	fmt.Printf("Go module '%s' generated successfully with dependencies: %v.\n", project, dependencies)
	return nil
}
