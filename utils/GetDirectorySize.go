package utils

import (
	"os"
	"path/filepath"
)

func GetDirectorySize(DirectoryPath string) (int64, string, error) {
	var size int64
	err := filepath.Walk(DirectoryPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, ConvertBytesToString(size), err
}
