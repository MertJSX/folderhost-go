package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src, dest string, cb func(int64, bool, string)) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("cannot open zip file: %v", err)
	}
	defer r.Close()

	err = os.MkdirAll(dest, 0777)

	if err != nil {
		return fmt.Errorf("cannot create folder: %v", err)
	}

	var (
		totalSize            int64 = 0
		remainingFolderSpace int64 = GetRemainingFolderSpace()
		currentUID           int   = os.Getuid()
		currentGID           int   = os.Getgid()
	)

	for _, file := range r.File {
		fmt.Printf("Process: %s\n", ConvertBytesToString(totalSize))
		cb(totalSize, false, "") // Parameters: totalSize, isCompleted, abortMsg
		err := extractFile(file, dest, &totalSize, currentUID, currentGID)
		if err != nil {
			return fmt.Errorf("unable to extract file (%s): %v", file.Name, err)
		}
		if totalSize > remainingFolderSpace {
			err := os.RemoveAll(dest)
			if err != nil {
				cb(totalSize, false, "Unzip process exceeds storage limit! Error while deleting the extracted folder.")
				return fmt.Errorf("unzip process exceeds storage limit")
			}
			cb(totalSize, false, "Unzip process exceeds storage limit!")
			return fmt.Errorf("unzip process exceeds storage limit")
		}
	}

	fmt.Println("Completed successfully!")
	cb(totalSize, true, "")

	return nil
}

func extractFile(file *zip.File, dest string, totalSize *int64, uid int, gid int) error {
	filePath := filepath.Join(dest, file.Name)

	if !IsSafePath(filePath) {
		return fmt.Errorf("security risk: wrong filepath")
	}

	if file.FileInfo().IsDir() {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			return err
		}
		return os.Chown(filePath, uid, gid)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	size, err := io.Copy(outFile, rc)
	*totalSize += size
	return err
}
