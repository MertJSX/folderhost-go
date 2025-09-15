package utils

import (
	"path/filepath"
	"strings"
)

func GetPureFileName(path string) string {
	fileName := filepath.Base(path)

	firstDot := strings.Index(fileName, ".")
	if firstDot > 0 {
		return fileName[:firstDot]
	}
	return fileName
}
