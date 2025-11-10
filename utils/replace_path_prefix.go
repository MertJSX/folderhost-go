package utils

import "strings"

func ReplacePathPrefix(fullPath string, realPrefix string) string {
	fullPath = strings.ReplaceAll(fullPath, "\\", "/")

	if strings.HasPrefix(fullPath, realPrefix) {
		return "./" + fullPath[len(realPrefix):]
	}

	return fullPath
}
