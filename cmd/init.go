package cmd

import (
	"einar/cmd/archetype"

	"github.com/spf13/cobra"
)

var moduleName string

var dependencies []string = []string{
	"github.com/joho/godotenv",
	"github.com/rs/zerolog/log",
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [module-name]",
	Short: "Initialize a new Go module",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	if err := archetype.CreateRootDirectory(moduleName); err != nil {
		return
	}
	if err := archetype.CreateMainFile(moduleName); err != nil {
		return
	}
	if err := archetype.InitializeGoModule(moduleName, dependencies); err != nil {
		return
	}
	if err := archetype.CreateArchetypeConfiguration(moduleName); err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&moduleName, "name", "n", "", "Name of the Go module")
	initCmd.MarkFlagRequired("name")
}
