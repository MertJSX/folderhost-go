package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/types"
)

func GetDirectoryItems(directoryPath string, mode string, config types.ConfigFile) []types.DirectoryItem {
	var directoryItems []types.DirectoryItem
	var id int = 0

	// Open the directory
	dir, err := os.Open(directoryPath)
	if err != nil {
		log.Printf("Error opening directory %s: %v", directoryPath, err)
		return nil
	}
	defer dir.Close()

	// Read only the immediate directory contents (non-recursive)
	files, err := dir.Readdir(-1) // -1 means return all entries
	if err != nil {
		log.Printf("Error reading directory %s: %v", directoryPath, err)
		return nil
	}

	for _, file := range files {
		fullPath := filepath.Join(directoryPath, file.Name())
		parentPath := GetParentPath(fullPath)

		fmt.Printf("Parent path: %s\n", parentPath)

		parentPath = ReplaceHostPrefix(parentPath, config)

		if parentPath[len(parentPath)-1] != '/' {
			parentPath += "/"
		}

		directoryItem := types.DirectoryItem{
			Id:           id,
			Name:         file.Name(),
			ParentPath:   parentPath,
			IsDirectory:  file.IsDir(),
			Path:         fmt.Sprintf("%s%s", parentPath, file.Name()),
			DateModified: file.ModTime(),
			Size:         ConvertBytesToString(file.Size()),
			SizeBytes:    file.Size(),
		}

		directoryItems = append(directoryItems, directoryItem)
		id++
	}

	return directoryItems
}
