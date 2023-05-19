package installations

import (
	"archetype/cmd/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func InstallPubSub(project string) error {
	// Read the JSON config file
	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Find the PubSub libraries
	var pubsubLibs []string
	for _, installCommand := range config.InstallationCommands {
		if installCommand.Name == "pubsub" {
			pubsubLibs = installCommand.Libraries
			break
		}
	}

	if len(pubsubLibs) == 0 {
		err = fmt.Errorf("pubsub libraries not found in .einar.cli.json")
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
	sourceDir := filepath.Join(filepath.Dir(binaryPath), "app/shared/archetype/pubsub")
	destDir := filepath.Join(project, "app/shared/archetype/pubsub")

	// Clone the source directory to the destination directory
	err = utils.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning pubsub directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("PubSub directory cloned successfully to %s.\n", destDir)

	configPath := filepath.Join(project, ".einar.cli.json")

	var wg sync.WaitGroup
	errChan := make(chan error, len(pubsubLibs))
	// Install pubsub libraries
	for _, lib := range pubsubLibs {
		wg.Add(1)
		go func(lib string) {
			defer wg.Done()

			cmd := exec.Command("go", "get", lib)
			cmd.Dir = project
			cmd.Stdout = os.Stdout // Command's stdout will be attached to system's stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				errChan <- fmt.Errorf("error installing pubsub library %s: %v", lib, err)
				return
			}

			// Add the installed library to the JSON config
			if err := AddInstallation(configPath, "pubsub", lib /*version*/, ""); err != nil {
				errChan <- fmt.Errorf("failed to update .einar.cli.latest.json: %v", err)
				return
			}
		}(lib)
	}

	wg.Wait() // Wait for all goroutines to finish
	close(errChan)

	// Check if any goroutine returned an error
	for err := range errChan {
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Update setup.go file with the import statement
	setupFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = utils.AddImportStatement(setupFilePath, "archetype/app/shared/archetype/pubsub")
	if err != nil {
		fmt.Println("Failed to add import statement to setup.go:", err)
		return err
	}

	return nil
}
