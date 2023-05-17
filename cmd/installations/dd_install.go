package installations

import "fmt"

func DDInstall(installation string) bool {
	// Continue with the installation
	switch installation {
	case "dd-chi-server":
		if err := InstallChiServer(""); err != nil {
			fmt.Println("Failed to install Chi Server:", err)
		} else {
			fmt.Println("Chi Server installed successfully.")
			return true
		}
	case "dd-pubsub":
		if err := InstallPubSub(""); err != nil {
			fmt.Println("Failed to install Pubsub Client:", err)
		} else {
			fmt.Println("Pubsub Client installed successfully.")
			return true
		}
	case "dd-firestore":
		if err := InstallFirestore(""); err != nil {
			fmt.Println("Failed to install Firestore Database:", err)
		} else {
			fmt.Println("Firestore Database installed successfully.")
			return true
		}
	case "dd-postgres":
		if err := InstallPostgres(""); err != nil {
			fmt.Println("Failed to install Postgres Database:", err)
		} else {
			fmt.Println("Postgres Database installed successfully.")
			return true
		}
	case "dd-redis":
		if err := InstallRedis(""); err != nil {
			fmt.Println("Failed to install Redis Database:", err)
		} else {
			fmt.Println("Redis Database installed successfully.")
			return true
		}
	case "dd-resty":
		if err := InstallResty(""); err != nil {
			fmt.Println("Failed to install Resty Client:", err)
		} else {
			fmt.Println("Resty Client installed successfully.")
			return true
		}
	}
	return false
}
