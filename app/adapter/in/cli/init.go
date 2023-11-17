package cli

import (
	"fmt"

	"github.com/Ignaciojeria/einar/app/business"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"github.com/Ignaciojeria/einar/app/shared/utils"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(initCmd)
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
	repositoryURL := args[1]
	utils.GitCloneTemplateInBinaryPath(repositoryURL, args[2])
	templatePath, err := utils.GetTemplateFolderPath(repositoryURL)
	if err != nil {
		fmt.Println("error getting template path")
		return
	}
	project := args[0]
	if args[0] == "." {
		project, _ = utils.GetCurrentFolderName()
	}
	project = utils.ConvertStringCase(project, "kebab")
	business.EinarInit(cmd.Context(), templatePath, project)
}
