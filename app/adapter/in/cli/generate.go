package cli

import (
	"fmt"

	"github.com/Ignaciojeria/einar/app/business"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"github.com/Ignaciojeria/einar/app/shared/utils"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(generateCmd)
}

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
	if err := business.EinarGenerate(
		cmd.Context(),
		config.Project,
		componentKind,
		componentName); err != nil {
		return
	}
	fmt.Println("Generate command executed for:", componentKind, "with name:", componentName)
}
