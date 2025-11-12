package utils

import (
	"github.com/MertJSX/folder-host-go/utils/config"
)

func GetRemainingFolderSpace() (int64, error) {
	mainFolderSize, _, err := GetDirectorySize(config.Config.Folder)

	if err != nil {
		return 0, err
	}

	fileCount := GetActiveFileCount()
	editorUsage := int64(fileCount * 200 * 1024)

	return config.Config.SizeBytes - (mainFolderSize + editorUsage), nil
}
