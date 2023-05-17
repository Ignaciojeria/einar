package utils

import "os"

func GetBinaryPath() string {
	binaryPath, _ := os.Executable()
	return binaryPath
}
