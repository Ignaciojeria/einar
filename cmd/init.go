package cmd

import (
	"archetype/cmd/base"
	"archetype/cmd/installations"
	"fmt"

	"github.com/spf13/cobra"
)

var project string

var dependency_tree []string = []string{
	"github.com/joho/godotenv@v1.5.1",
	"github.com/rs/zerolog@v1.29.1",
	"github.com/go-chi/chi@v1.5.4",
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go module",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {

	config, err := base.ReadEinarCli()
	if err != nil {
		fmt.Println(err)
		return
	}
	dependency_tree = make([]string, 0)
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
		fmt.Println("Specify a subcommand")
	},
}

var initMode string

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&project, "name", "n", "", "Name of the project")
	initCmd.Flags().StringVarP(&initMode, "mode", "m", "all-in-one", "Mode of initialization (default or all-in-one)")
	initCmd.MarkFlagRequired("name")
}
