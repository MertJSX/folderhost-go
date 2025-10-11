package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/resources"
)

func Setup() {
	if IsNotExistingPath("tmp") {
		fmt.Println("Creating /tmp folder...")
		err := os.Mkdir("tmp", 0700)

		if err != nil {
			log.Fatalf("Error creating tmp folder!")
		}
	} else {
		os.RemoveAll("tmp")
		os.Mkdir("tmp", 0700)
	}

	if IsNotExistingPath("./config.yml") {
		fmt.Println("Creating config file...")
		configContent, err := resources.DefaultConfig.ReadFile("default_config.yml")

		if err != nil {
			log.Fatalf("Error reading embedded file: %s", err)
		}

		err = os.WriteFile("config.yml", configContent, 0700)

		if err != nil {
			log.Fatalf("Error creating config.yml")
		}
	}

	if IsNotExistingPath("recovery_bin") && Config.RecoveryBin {
		fmt.Println("Creating /recovery_bin folder...")
		err := os.Mkdir("recovery_bin", 0700)

		if err != nil {
			log.Fatalf("Error creating recovery_bin folder!")
		}
	}
}

func AutoClearOldLogs() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	if Config.ClearLogsAfter > 0 {
		err := logs.ClearOldLogs(Config.ClearLogsAfter)
		if err != nil {
			fmt.Printf("Error while clearing old logs: %s\n", err)
		}
	} else {
		return
	}

	for range ticker.C {
		if Config.ClearLogsAfter > 0 {
			err := logs.ClearOldLogs(Config.ClearLogsAfter)
			if err != nil {
				fmt.Printf("Error while clearing old logs: %s\n", err)
			}
		}
	}
}
