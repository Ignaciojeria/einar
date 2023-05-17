package installations

import (
	"archetype/cmd/base"
	"archetype/cmd/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallFirestore(project string) error {
	// Read the JSON config file
	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the Firestore libraries
	var firestoreLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "firestore" {
			firestoreLibs = installCommand.Libraries
			break
		}
	}

	if len(firestoreLibs) == 0 {
		err = fmt.Errorf("firestore libraries not found in .einar.cli.json")
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
	sourceDir := filepath.Join(filepath.Dir(binaryPath), "app/shared/archetype/firestore")
	destDir := filepath.Join(project, "app/shared/archetype/firestore")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning firestore directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("firestore directory cloned successfully to %s.\n", destDir)

	configPath := filepath.Join(project, ".einar.cli.json")

	// Install Firestore libraries
	for _, lib := range firestoreLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing firestore library %s: %v", lib, err)
			fmt.Println(err)
			return err
		}

		// Add the installed library to the JSON config
		if err := AddInstallation(configPath, "firestore", lib /*version*/, ""); err != nil {
			fmt.Println("Failed to update .einar.cli.latest.json:", err)
			return err
		}
	}

	// Update setup.go file with the import statement
	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = utils.AddImportStatement(setupFilePath, "archetype/app/shared/archetype/firestore")
	if err != nil {
		fmt.Println("Failed to add import statement to setup.go:", err)
		return err
	}

	return nil
}
