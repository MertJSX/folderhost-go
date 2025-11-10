package utils

import "strings"

func GetParentPath(DirectoryItemPath string) string {
	DirectoryItemPath = strings.ReplaceAll(DirectoryItemPath, "\\", "/")
	lastIndex := strings.LastIndex(DirectoryItemPath, "/")

	if lastIndex == -1 {
		return DirectoryItemPath
	}

	item := DirectoryItemPath[0:lastIndex]

	if len(item) > 1 {
		return item
	} else {
		return DirectoryItemPath[0 : lastIndex+1]
	}
}
