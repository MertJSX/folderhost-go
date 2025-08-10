package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/MertJSX/folder-host-go/types"
)

type directoryID struct {
	id       int
	fullPath string
}

func GetDirectoryItems(directoryPath string, mode string, config types.ConfigFile) ([]types.DirectoryItem, int64) {
	var directoryItems []types.DirectoryItem
	var directoryIDs []directoryID

	fmt.Printf("Mode: %s\n", mode)

	// Open the directory
	dir, err := os.Open(directoryPath)
	if err != nil {
		log.Printf("Error opening directory %s: %v", directoryPath, err)
		return nil, 0
	}
	defer dir.Close()

	// Read only the immediate directory contents (non-recursive)
	files, err := dir.Readdir(-1) // -1 means return all entries
	if err != nil {
		log.Printf("Error reading directory %s: %v", directoryPath, err)
		return nil, 0
	}

	for i, file := range files {
		fullPath := filepath.Join(directoryPath, file.Name())
		parentPath := GetParentPath(fullPath)
		parentPath = ReplaceHostPrefix(parentPath, config)
		if parentPath[len(parentPath)-1] != '/' {
			parentPath += "/"
		}
		path := fmt.Sprintf("%s%s", parentPath, file.Name())
		isDir := file.IsDir()
		size := ConvertBytesToString(file.Size())
		var sizeBytes int64 = file.Size()

		if isDir && mode == "Quality mode" {
			directoryIDs = append(directoryIDs, directoryID{id: i, fullPath: fullPath})
		}

		directoryItem := types.DirectoryItem{
			Id:           i,
			Name:         file.Name(),
			ParentPath:   parentPath,
			IsDirectory:  isDir,
			Path:         path,
			DateModified: file.ModTime(),
			Size:         size,
			SizeBytes:    sizeBytes,
		}

		directoryItems = append(directoryItems, directoryItem)
	}

	if mode == "Quality mode" {
		ch := make(chan types.DirectorySizeOutput, len(directoryIDs))
		var wg sync.WaitGroup
		for _, dirID := range directoryIDs {
			wg.Add(1)
			go GetDirectorySizeAsync(dirID.fullPath, dirID.id, ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for result := range ch {
			fmt.Printf("Path: %s\n", directoryItems[result.Id].Path)
			fmt.Printf("Size: %s\n", result.Size)
			directoryItems[result.Id].Size = result.Size
			directoryItems[result.Id].SizeBytes = result.SizeBytes
		}
	}

	var totalSize int64 = 0

	if mode != "Optimized mode" {
		for _, dirItem := range directoryItems {
			totalSize += dirItem.SizeBytes
		}
	}

	return directoryItems, totalSize
}
