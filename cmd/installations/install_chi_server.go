package installations

import (
	"archetype/cmd/base"
	"archetype/cmd/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallChiServer(project string) error {
	// Read the JSON config file
	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the chi-server libraries
	var chiLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "chi-server" {
			chiLibs = installCommand.Libraries
			break
		}
	}

	if len(chiLibs) == 0 {
		err = fmt.Errorf("chi-server libraries not found in .einar.cli.json")
		fmt.Println(err)
		return err
	}

	// Get the binary path
	binaryPath, err := os.Executable()
	if err != nil {
		err = fmt.Errorf("failed to get binary path: %v", err)
		fmt.Println(err)
		return err
	}

	// Define the source and destination directories
	sourceDir := filepath.Join(filepath.Dir(binaryPath), "app/shared/archetype/chi_server")
	destDir := filepath.Join(project, "app/shared/archetype/chi_server")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning chi_server directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("chi_server directory cloned successfully to %s.\n", destDir)

	configPath := filepath.Join(project, ".einar.cli.json")

	// Install chi-server libraries
	for _, lib := range chiLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing chi-server library %s: %v", lib, err)
			fmt.Println(err)
			return err
		}

		// Add the installed library to the JSON config
		if err := AddInstallation(configPath, "chi-server", lib /*version*/, ""); err != nil {
			fmt.Println("Failed to update .einar.cli.latest.json:", err)
			return err
		}
	}

	// Update setup.go file with the import statement
	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = utils.AddImportStatement(setupFilePath, "archetype/app/shared/archetype/chi_server")
	if err != nil {
		fmt.Println("Failed to add import statement to setup.go:", err)
		return err
	}

	return nil
}
