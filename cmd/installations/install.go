package installations

import "fmt"

func Install(installationCommand string) bool {
	// Continue with the installation
	if err := InstallCommand("", installationCommand); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
