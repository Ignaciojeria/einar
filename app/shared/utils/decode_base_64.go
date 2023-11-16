package utils

import (
	"encoding/base64"

	"errors"
)

func DecodeBase64(b64 string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		b, err = base64.StdEncoding.DecodeString(b64 + "=")
	}
	if err != nil {
		b, err = base64.StdEncoding.DecodeString(b64 + "==")
	}
	if err != nil {
		return nil, errors.New("error decoding b64 : " + b64)
	}
	return b, nil
}
