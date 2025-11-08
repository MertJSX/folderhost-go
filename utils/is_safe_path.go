package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/MertJSX/folder-host-go/utils/config"
)

func IsSafePath(requestedPath string) bool {
	baseDir := filepath.Clean(fmt.Sprintf("./%s", config.Config.Folder))

	fullPath := filepath.Join(baseDir, requestedPath)

	cleanPath := filepath.Clean(fullPath)

	return strings.HasPrefix(cleanPath, baseDir)
}
