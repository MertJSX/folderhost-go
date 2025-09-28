package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ClearDirectory(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("folder not found: %s", dirPath)
	}

	if !info.IsDir() {
		return fmt.Errorf("not a directory: %s", dirPath)
	}

	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	entries, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry)
		err = os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}
