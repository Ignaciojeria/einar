package utils

import (
	"errors"
	"strings"
)

// SplitCredentials splits a "username:password" string into two separate strings.
func SplitCredentials(credentials string) (string, string, error) {
	split := strings.Split(credentials, ":")

	if len(split) != 2 {
		return "", "", errors.New("invalid credentials format, expected 'username:password'")
	}

	return split[0], split[1], nil
}
