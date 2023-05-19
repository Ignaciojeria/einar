package utils

import "os"

func IsEinarCliJsonExists() bool {
	if _, err := os.Stat("einar.cli.json"); os.IsNotExist(err) {
		return false
	}
	return true
}
