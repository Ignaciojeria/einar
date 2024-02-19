package utils

import (
	"errors"
	"path"
)

func GetLatestTag(templatePath string) (string, error) {
	tag := path.Base(templatePath)
	if tag == "" || tag == "/" {
		return "", errors.New("no tag found in the URL")
	}

	return tag, nil
}
