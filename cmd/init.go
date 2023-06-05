package cmd

import (
	"archetype/cmd/base"
	"archetype/cmd/components"
	"archetype/cmd/installations"
	"archetype/cmd/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var repositoryURL string

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

	if repositoryURL == "" {
		repositoryURL = "https://github.com/Ignaciojeria/einar-cli-template" // Default repository URL
	}

	utils.GitCloneTemplateInBinaryPath(repositoryURL)

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

	dependencyTree := make([]string, 0)

	for _, installBase := range template.InstallationsBase {
		dependencyTree = append(dependencyTree, installBase.Library)
	}

	if err := base.InitializeGoModule(dependencyTree); err != nil {
		return
	}
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install command for Einar",
	Long:  `This command allows you to install various components.`,
	Args:  cobra.ExactArgs(1),
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

		if err := installations.InstallCommand("", args[0]); err != nil {
			return
		}
	},
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [component type] [component name]",
	Short: "generate component. for example: einar generate subscription my-subscription",
	Args:  cobra.ExactArgs(2), // Ensure exactly 2 arguments are provided
	Run:   runGenerateCmd,
}

func runGenerateCmd(cmd *cobra.Command, args []string) {
	componentKind := args[0]
	componentName := args[1]

	componentName = utils.ConvertStringCase(componentName, "kebab")

	if err := components.GenerateComponenteCommand(
		"",
		componentKind, componentName); err != nil {
		return
	}

	fmt.Println("Generate command executed for:", componentKind, "with name:", componentName)
}

func init() {
	initCmd.Flags().StringVar(&repositoryURL, "repository", "", "URL of the repository") // Add a flag for repository URL
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(generateCmd)
}
