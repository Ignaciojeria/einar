package cli

import (
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "short description of your command",
	Run:   runconnect,
}

func runconnect(cmd *cobra.Command, args []string) {
	fmt.Println("connect command not implemented yet")
}
