package cmd

import (
	"archetype/cmd/base"
	"archetype/cmd/installations"
	"archetype/cmd/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go module",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	_, err := utils.ReadEinarCli()
	if err == nil {
		fmt.Println("einar cli already initialized")
		return
	}
	utils.GitCloneTemplateInBinaryPath("https://github.com/Ignaciojeria/einar-cli-template")

	project, _ := utils.GetCurrentFolderName()

	if err := base.CreateFilesFromTemplate(project); err != nil {
		fmt.Println(err)
		return
	}

	if err := base.CreateDirectoriesFromTemplate(project); err != nil {
		fmt.Println(err)
		return
	}

	template, err := utils.ReadEinarTemplateFromBinaryPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	dependency_tree := make([]string, 0)

	for _, installBase := range template.InstallationsBase {
		dependency_tree = append(dependency_tree, installBase.Library)
	}

	if err := base.InitializeGoModule(dependency_tree); err != nil {
		return
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

		fmt.Println("Unknown installation command.")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(installCmd)
}
