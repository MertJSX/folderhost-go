package tasks

import (
	"log"
	"time"

	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
)

func UpdateRemainingFolderSpace() {
	remainingFS, err := utils.GetRemainingFolderSpace()

	if err != nil {
		log.Fatalf("Error updating remaining folder space: %v\n", err)
	}

	cache.RemainingFolderSpace = remainingFS

	time.Sleep(time.Minute * 2)
}
