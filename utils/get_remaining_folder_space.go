package utils

func GetRemainingFolderSpace() (int64, error) {
	mainFolderSize, _, err := GetDirectorySize(Config.Folder)

	if err != nil {
		return 0, err
	}

	fileCount := GetActiveFileCount()
	editorUsage := int64(fileCount * 200 * 1024)

	// fmt.Printf("Editor usage: %d KB\n", editorUsage/1024)

	return Config.SizeBytes - (mainFolderSize + editorUsage), nil
}
