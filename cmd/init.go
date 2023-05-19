package cmd

import (
	"archetype/cmd/base"
	"archetype/cmd/installations"
	"archetype/cmd/utils"
	"fmt"
	"strings"

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

	config, err := utils.ReadEinarCliFromBinaryPath()
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

		if config.IsInstalled(strings.ReplaceAll(args[0], "dd-", "")) {
			fmt.Println("installation " + args[0] + " already added")
			return
		}

		if installations.Install(args[0]) {
			return
		}

		if installations.DDInstall(args[0]) {
			return
		}

		fmt.Println("Unknown installation command.")

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
