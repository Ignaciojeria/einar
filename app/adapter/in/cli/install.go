package cli

import (
	"fmt"

	"github.com/Ignaciojeria/einar/app/business"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"github.com/Ignaciojeria/einar/app/shared/utils"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install command for Einar",
	Long:  `This command allows you to install various components.`,
	Args:  cobra.ExactArgs(1),
	Run:   runinstall,
}

func runinstall(cmd *cobra.Command, args []string) {
	config, _ := utils.ReadEinarCli()
	if config.Project == "${project}" {
		fmt.Println("Run installation command only inside your project.")
		return
	}

	if config.IsInstalled(args[0]) {
		fmt.Println("installation " + args[0] + " already added")
		return
	}

	if err := business.EinarInstall(cmd.Context(), config.Project, args[0]); err != nil {
		fmt.Println(err.Error())
		return
	}
}
