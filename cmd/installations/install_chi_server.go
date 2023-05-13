package installations

import (
	"archetype/cmd/base"
	"fmt"
	"os/exec"
	"path/filepath"
)

func InstallChiServer(project string) error {
	// Read the JSON config file
	config, err := base.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the resty libraries
	var chiLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "chi-server" {
			chiLibs = installCommand.Libraries
			break
		}
	}

	if len(chiLibs) == 0 {
		err = fmt.Errorf("chi-server libraries not found in .einar.cli.latest.json")
		fmt.Println(err)
		return err
	}

	// Define the source and destination directories
	sourceDir := "app/shared/archetype/chi_server"
	destDir := filepath.Join(project, "app/shared/archetype/chi_server")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning chi_server directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("chi_server directory cloned successfully to %s.\n", destDir)

	// Install resty libraries
	for _, lib := range chiLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing chi-server library %s: %v\n", lib, err)
			fmt.Println(err)
			return err
		}
	}

	return nil
}
