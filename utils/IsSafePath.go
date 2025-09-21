package utils

import "path/filepath"

func IsSafePath(basePath, targetPath string) bool {
	rel, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return false
	}

	if rel == ".." || len(rel) >= 2 && rel[:2] == ".." {
		return false
	}

	return !filepath.IsAbs(rel)
}
