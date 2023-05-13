package installations

import (
	"archetype/cmd/base"
	"fmt"
	"os/exec"
	"path/filepath"
)

func InstallPubSub(project string) error {
	// Read the JSON config file
	config, err := base.ReadEinarCli()
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
		err = fmt.Errorf("pubsub libraries not found in .einar.cli.latest.json")
		fmt.Println(err)
		return err
	}

	// Define the source and destination directories
	sourceDir := "app/shared/archetype/pubsub"
	destDir := filepath.Join(project, "app/shared/archetype/pubsub")

	// Clone the source directory to the destination directory
	err = base.CopyDirectory(sourceDir, destDir, project)
	if err != nil {
		err := fmt.Errorf("error cloning pubsub directory: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("PubSub directory cloned successfully to %s.\n", destDir)

	// Install pubsub libraries
	for _, lib := range pubsubLibs {
		cmd := exec.Command("go", "get", lib)
		cmd.Dir = project
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("error installing pubsub library %s: %v\n", lib, err)
			fmt.Println(err)
			return err
		}
	}

	return nil
}
