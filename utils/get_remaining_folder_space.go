package utils

func GetRemainingFolderSpace() (int64, error) {
	mainFolderSize, _, err := GetDirectorySize(Config.Folder)

	if err != nil {
		return 0, err
	}

	return Config.SizeBytes - mainFolderSize, nil
}
