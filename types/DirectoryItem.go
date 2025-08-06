package types

import "time"

type DirectoryItem struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	ParentPath   string    `json:"parentPath"`
	Path         string    `json:"path"`
	IsDirectory  bool      `json:"isDirectory"`
	DateModified time.Time `json:"dateModified"`
	Size         string    `json:"size"`
	SizeBytes    int64     `json:"sizeBytes"`
	StorageLimit string    `json:"storage_limit"`
}
