package installations

import (
	"archetype/cmd/base"
	"archetype/cmd/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallPostgres(project string) error {
	// Read the JSON config file
	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the postgres libraries
	var postgresLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "postgres" {
			postgresLibs = installCommand.Libraries
			break
		}
	}

	if len(postgresLibs) == 0 {
		err = fmt.Errorf("postgres libraries not found in .einar.cli.json")
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
	sourceDir := filepath.Join(filepath.Dir(binaryPath), "app/shared/archetype/postgres")
	destDir := filepath.Join(project, "app/shared/archetype/postgres")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning postgres directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("postgres directory cloned successfully to %s.\n", destDir)

	configPath := filepath.Join(project, ".einar.cli.json")

	// Install postgres libraries
	for _, lib := range postgresLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing postgres library %s: %v", lib, err)
			fmt.Println(err)
			return err
		}

		// Add the installed library to the JSON config
		if err := AddInstallation(configPath, "postgres", lib /*version*/, ""); err != nil {
			fmt.Println("Failed to update .einar.cli.latest.json:", err)
			return err
		}
	}

	// Update setup.go file with the import statement
	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = utils.AddImportStatement(setupFilePath, "archetype/app/shared/archetype/postgres")
	if err != nil {
		fmt.Println("Failed to add import statement to setup.go:", err)
		return err
	}

	return nil
}
