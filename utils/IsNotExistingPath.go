package utils

import "os"

func IsNotExistingPath(path string) bool {
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}
