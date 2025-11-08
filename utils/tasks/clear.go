package tasks

import (
	"fmt"
	"time"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/utils/config"
)

func AutoClearOldLogs() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	config := &config.Config

	if config.ClearLogsAfter > 0 {
		err := logs.ClearOldLogs(config.ClearLogsAfter)
		if err != nil {
			fmt.Printf("Error while clearing old logs: %s\n", err)
		}
	} else {
		return
	}

	for range ticker.C {
		if config.ClearLogsAfter > 0 {
			err := logs.ClearOldLogs(config.ClearLogsAfter)
			if err != nil {
				fmt.Printf("Error while clearing old logs: %s\n", err)
			}
		}
	}
}
