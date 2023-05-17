package cmd

import (
	"archetype/cmd/base"
	"archetype/cmd/installations"
	"archetype/cmd/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var project string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go module",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {

	config, err := utils.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return
	}
	dependency_tree := make([]string, 0)
	for _, installBase := range config.InstallationsBase {
		dependency_tree = append(dependency_tree, installBase.Library)
	}

	if err := base.CreateRootDirectory(project); err != nil {
		return
	}
	if err := base.CreateMainFile(project); err != nil {
		return
	}
	if err := base.InitializeGoModule(project, dependency_tree); err != nil {
		return
	}
	if err := base.CreateConfiguration(project); err != nil {
		return
	}
	if err := base.CreateContainer(project); err != nil {
		return
	}
	if err := base.CreateEnvironment(project); err != nil {
		return
	}
	if err := base.CreateGitignore(project); err != nil {
		return
	}
	if err := base.CreateVersion(project); err != nil {
		return
	}
	if err := base.CreateUtils(project); err != nil {
		return
	}
	if err := base.CreateArchetypeSetupFile(project); err != nil {
		return
	}
	if err := base.CreateEinarCli(project); err != nil {
		return
	}

	if initMode == "all-in-one" {
		if err := installations.InstallResty(project); err != nil {
			return
		}
		if err := installations.InstallPubSub(project); err != nil {
			return
		}
		if err := installations.InstallChiServer(project); err != nil {
			return
		}
		if err := installations.InstallPostgres(project); err != nil {
			return
		}
		if err := installations.InstallRedis(project); err != nil {
			return
		}
		if err := installations.InstallFirestore(project); err != nil {
			return
		}
	}
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install command for Einar",
	Long:  `This command allows you to install various components.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Read the JSON config file
		config, _ := utils.ReadEinarCli()
		if config.Project == "${project}" {
			fmt.Println("Run installation command only inside your project.")
			return
		}
		if config.IsInstalled(args[0]) {
			fmt.Println("installation " + args[0] + " already added")
			return
		}
		// Continue with the installation
		switch args[0] {
		case "chi-server":
			if err := installations.InstallChiServer(""); err != nil {
				fmt.Println("Failed to install Chi Server:", err)
			} else {
				fmt.Println("Chi Server installed successfully.")
			}
		case "pubsub":
			if err := installations.InstallPubSub(""); err != nil {
				fmt.Println("Failed to install Pubsub Client:", err)
			} else {
				fmt.Println("Pubsub Client installed successfully.")
			}
		case "firestore":
			if err := installations.InstallFirestore(""); err != nil {
				fmt.Println("Failed to install Firestore Database:", err)
			} else {
				fmt.Println("Firestore Database installed successfully.")
			}
		case "postgres":
			if err := installations.InstallPostgres(""); err != nil {
				fmt.Println("Failed to install Postgres Database:", err)
			} else {
				fmt.Println("Postgres Database installed successfully.")
			}
		case "redis":
			if err := installations.InstallRedis(""); err != nil {
				fmt.Println("Failed to install Redis Database:", err)
			} else {
				fmt.Println("Redis Database installed successfully.")
			}
		case "resty":
			if err := installations.InstallResty(""); err != nil {
				fmt.Println("Failed to install Resty Client:", err)
			} else {
				fmt.Println("Resty Client installed successfully.")
			}
		default:
			fmt.Println("Unknown installation command.")
		}
	},
}

var initMode string

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&project, "name", "n", "", "Name of the project")
	initCmd.Flags().StringVarP(&initMode, "mode", "m", "all-in-one", "Mode of initialization (default or all-in-one)")
	initCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(installCmd)
}
