package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/types"
)

func GetDirectoryItems(DirectoryPath string, Mode string, Config types.ConfigFile) {
	var directoryItems []types.DirectoryItem
	var id int

	err := filepath.Walk(DirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("-----------------")
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Parent Path: %s\n", GetParentPath(path))
		fmt.Printf("Name: %s\n", info.Name())
		fmt.Printf("IsDirectory %t\n", info.IsDir())
		fmt.Printf("Size %s\n", ConvertBytesToString(info.Size()))
		fmt.Printf("Date modified: %s\n", info.ModTime())
		fmt.Println("-----------------")
		var DirectoryItem types.DirectoryItem = types.DirectoryItem{
			Id:           id,
			Name:         info.Name(),
			ParentPath:   GetParentPath(path),
			IsDirectory:  info.IsDir(),
			Path:         path,
			DateModified: info.ModTime(),
			Size:         ConvertBytesToString(info.Size()),
			SizeBytes:    info.Size(),
		}

		directoryItems = append(directoryItems, DirectoryItem)

		id++

		return nil
	})
	if err != nil {
		log.Fatalf("Error while Getting directory items of path %s", DirectoryPath)
	}

}
