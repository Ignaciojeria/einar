package installations

import "fmt"

func Install(installation string) bool {
	// Continue with the installation
	switch installation {
	case "chi-server":
		if err := InstallChiServer(""); err != nil {
			fmt.Println("Failed to install Chi Server:", err)
		} else {
			fmt.Println("Chi Server installed successfully.")
			return true
		}
	case "pubsub":
		if err := InstallPubSub(""); err != nil {
			fmt.Println("Failed to install Pubsub Client:", err)
		} else {
			fmt.Println("Pubsub Client installed successfully.")
			return true
		}
	case "firestore":
		if err := InstallFirestore(""); err != nil {
			fmt.Println("Failed to install Firestore Database:", err)
		} else {
			fmt.Println("Firestore Database installed successfully.")
			return true
		}
	case "postgres":
		if err := InstallPostgres(""); err != nil {
			fmt.Println("Failed to install Postgres Database:", err)
		} else {
			fmt.Println("Postgres Database installed successfully.")
			return true
		}
	case "redis":
		if err := InstallRedis(""); err != nil {
			fmt.Println("Failed to install Redis Database:", err)
		} else {
			fmt.Println("Redis Database installed successfully.")
			return true
		}
	case "resty":
		if err := InstallResty(""); err != nil {
			fmt.Println("Failed to install Resty Client:", err)
		} else {
			fmt.Println("Resty Client installed successfully.")
			return true
		}
	default:
		return false
	}
	return false
}
