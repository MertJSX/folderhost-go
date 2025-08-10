package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/MertJSX/folder-host-go/types"
)

func GetDirectorySize(DirectoryPath string) (int64, string, error) {
	var size int64
	start := time.Now()
	err := filepath.Walk(DirectoryPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	fmt.Printf("%s walked in %s\n", DirectoryPath, time.Since(start))
	return size, ConvertBytesToString(size), err
}

func GetDirectorySizeAsync(DirectoryPath string, id int, ch chan<- types.DirectorySizeOutput, wg *sync.WaitGroup) {
	defer wg.Done()

	size, sizeStr, err := GetDirectorySize(DirectoryPath)

	ch <- types.DirectorySizeOutput{
		SizeBytes: size,
		Size:      sizeStr,
		Id:        id,
		Error:     err,
	}
}
