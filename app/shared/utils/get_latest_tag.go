package utils

import (
	"errors"
	"path/filepath"
)

func GetLatestTag(templatePath string) (string, error) {
	if templatePath == "" {
		return "", errors.New("path is empty")
	}
	tag := filepath.Base(templatePath)
	if tag == "." || tag == "\\" || tag == "/" {
		return "", errors.New("no tag found in the path")
	}
	return tag, nil
}
