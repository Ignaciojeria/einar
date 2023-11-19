package cli

import (
	"fmt"

	"github.com/Ignaciojeria/einar/app/shared/archetype"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"

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
	if err := archetype.Setup(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect command not implemented yet")
}
