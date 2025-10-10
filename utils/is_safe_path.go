package utils

import "strings"

func IsSafePath(filePath string) bool {
	return !strings.Contains(filePath, "..")
}
