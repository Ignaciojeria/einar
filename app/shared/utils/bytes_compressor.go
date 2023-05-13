package utils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
)

func CompressBytes(s []byte) []byte {
	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	zipped.Write(s)
	zipped.Close()
	return zipbuf.Bytes()
}

func DecompressBytes(s []byte) ([]byte, error) {
	rdr, err := gzip.NewReader(bytes.NewReader(s))

	customError := errors.New("error decompresing bytes")

	if err != nil {
		return nil, customError
	}

	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, customError
	}
	rdr.Close()
	return data, nil
}
