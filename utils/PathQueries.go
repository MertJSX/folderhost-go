package utils

import "os"

func IsNotExistingPath(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}

func IsExistingPath(path string) bool {
	return !IsNotExistingPath(path)
}
