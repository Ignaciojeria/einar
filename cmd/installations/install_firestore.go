package installations

import (
	"archetype/cmd/base"
	"fmt"
	"os/exec"
	"path/filepath"
)

func InstallFirestore(project string) error {
	// Read the JSON config file
	config, err := base.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the resty libraries
	var chiLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "firestore" {
			chiLibs = installCommand.Libraries
			break
		}
	}

	if len(chiLibs) == 0 {
		err = fmt.Errorf("firestore libraries not found in .einar.cli.latest.json")
		fmt.Println(err)
		return err
	}

	// Define the source and destination directories
	sourceDir := "app/shared/archetype/firestore"
	destDir := filepath.Join(project, "app/shared/archetype/firestore")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning firestore directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("firestore directory cloned successfully to %s.\n", destDir)

	// Install resty libraries
	for _, lib := range chiLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing firestore library %s: %v\n", lib, err)
			fmt.Println(err)
			return err
		}
	}

	return nil
}
