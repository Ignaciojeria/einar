package installations

import (
	"archetype/cmd/base"
	"archetype/cmd/utils"
	"fmt"
	"os/exec"
	"path/filepath"
)

func InstallResty(project string) error {
	// Read the JSON config file
	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the resty libraries
	var restyLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "resty" {
			restyLibs = installCommand.Libraries
			break
		}
	}

	if len(restyLibs) == 0 {
		err = fmt.Errorf("resty libraries not found in .einar.cli.json")
		fmt.Println(err)
		return err
	}

	// Define the source and destination directories
	sourceDir := "app/shared/archetype/resty"
	destDir := filepath.Join(project, "app/shared/archetype/resty")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning resty directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Resty directory cloned successfully to %s.\n", destDir)

	// Install resty libraries
	for _, lib := range restyLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing resty library %s: %v\n", lib, err)
			fmt.Println(err)
			return err
		}
	}

	return nil
}
