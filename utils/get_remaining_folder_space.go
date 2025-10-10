package utils

func GetRemainingFolderSpace() int64 {
	maxSize := ConvertStringToBytes(Config.StorageLimit)
	mainFolderSize, _, _ := GetDirectorySize(Config.Folder)
	return maxSize - mainFolderSize
}
