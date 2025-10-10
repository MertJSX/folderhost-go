package utils

import (
	"os"
	"path/filepath"
)

func IsDevelopment() bool {
	if _, err := os.Stat("client/dist"); err == nil {
		return false // that's for testing client build on server
	}

	if len(os.Args) > 0 && filepath.Base(os.Args[0]) == "main" {
		return true
	}

	if _, err := os.Stat("client/src"); err == nil {
		return true
	}

	return false
}
