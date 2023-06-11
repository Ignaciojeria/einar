package cmd

import (
	"github.com/Ignaciojeria/einar-cli/einar/cmd/base"
	"github.com/Ignaciojeria/einar-cli/einar/cmd/components"
	"github.com/Ignaciojeria/einar-cli/einar/cmd/installations"
	"github.com/Ignaciojeria/einar-cli/einar/cmd/utils"
	"fmt"
	"github.com/spf13/cobra"
)
// initCmd represents the init command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "retrieve einar version",
	Run:   runVersionCmd,
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("beta")
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project name] [repository template]",
	Short: "Initialize a new Go module",
	Args:  cobra.ExactArgs(3),
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {

	_, err := utils.ReadEinarCli()

	if err == nil {
		fmt.Println("einar cli already initialized")
		return
	}

	utils.GitCloneTemplateInBinaryPath(args[1], args[2])

	project := args[0]
	if args[0] == "." {
		project, _ = utils.GetCurrentFolderName()
	}
	project = utils.ConvertStringCase(project, "kebab")
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

	if err := base.InitializeGoModule(dependencyTree, project); err != nil {
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

		if err := installations.InstallCommand(config.Project, args[0]); err != nil {
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

	// Read the JSON config file
	config, _ := utils.ReadEinarCli()
	if config.Project == "${project}" {
		fmt.Println("Run installation command only inside your project.")
		return
	}

	if err := components.GenerateComponenteCommand(
		config.Project,
		componentKind, componentName); err != nil {
		return
	}

	fmt.Println("Generate command executed for:", componentKind, "with name:", componentName)
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(versionCmd)
}
