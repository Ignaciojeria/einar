package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var moduleName string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [module-name]",
	Short: "Initialize a new Go module",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	// Ensure the module name is valid
	if moduleName == "" {
		fmt.Println("Please specify a name for the new module")
		return
	}

	// Create the module directory and the main.go file
	err := os.Mkdir(moduleName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory for module %s: %s\n", moduleName, err)
		return
	}
	mainFile, err := os.Create(filepath.Join(moduleName, "main.go"))
	if err != nil {
		fmt.Printf("Error creating main.go file for module %s: %s\n", moduleName, err)
		return
	}
	mainFile.Close()

	// Initialize a new Go module
	goModCmd := exec.Command("go", "mod", "init", moduleName)
	goModCmd.Dir = moduleName
	err = goModCmd.Run()
	if err != nil {
		fmt.Printf("Error initializing Go module for module %s: %s\n", moduleName, err)
		return
	}

	// Print success message
	fmt.Printf("Go module '%s' generated successfully.\n", moduleName)
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&moduleName, "name", "n", "", "Name of the Go module")
	initCmd.MarkFlagRequired("name")
}
