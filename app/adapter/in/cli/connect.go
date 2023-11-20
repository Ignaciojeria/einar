package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Ignaciojeria/einar/app/shared/archetype"
	"github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "short description of your command",
	Run:   runconnect,
}

func init() {
	cmd.RootCmd.AddCommand(connectCmd)
}

func runconnect(cmd *cobra.Command, args []string) {
	if len(args) > 0 && args[0] == "setup-child" {
		// This is the child process for setup
		if err := archetype.Setup(); err != nil {
			fmt.Fprintf(os.Stderr, "Setup failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Setup completed")
		os.Exit(0)
	}

	// Start the setup in a separate child process
	childProcess := exec.Command(os.Args[0], "connect", "setup-child")
	childProcess.Stdout = os.Stdout
	childProcess.Stderr = os.Stderr

	if err := childProcess.Start(); err != nil {
		fmt.Printf("Error starting setup child process: %s\n", err)
		return
	}

	fmt.Printf("Setup child process started, PID: %d\n", childProcess.Process.Pid)
	// You can decide to wait or not for the child process
	// err := childProcess.Wait()
	// if err != nil {
	//     fmt.Printf("Setup child process exited with error: %s\n", err)
	// }
}
