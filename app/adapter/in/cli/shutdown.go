package cli

import (
	"github.com/Ignaciojeria/einar/app/adapter/out/client"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(shutdownCmd)
}

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "short description of your command",
	Run:   runshutdown,
}

func runshutdown(cmd *cobra.Command, args []string) {
	client.Shutdown(cmd.Context())
}
