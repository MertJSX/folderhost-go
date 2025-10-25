package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func IsSafePath(requestedPath string) bool {
	baseDir := filepath.Clean(fmt.Sprintf("./%s", Config.Folder))

	fullPath := filepath.Join(baseDir, requestedPath)

	cleanPath := filepath.Clean(fullPath)

	return strings.HasPrefix(cleanPath, baseDir)
}
