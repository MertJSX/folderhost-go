package utils

import "strings"

func ReplacePathPrefix(fullPath string, realPrefix string) string {
	if strings.HasPrefix(fullPath, realPrefix) {
		return "./" + fullPath[len(realPrefix):]
	}

	return fullPath
}
