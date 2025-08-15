package utils

import (
	"bytes"
	"mime/multipart"
)

func FileToString(file multipart.File) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
