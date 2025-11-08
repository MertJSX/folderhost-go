package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/utils/config"
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

	if IsNotExistingPath("recovery_bin") && config.Config.RecoveryBin {
		fmt.Println("Creating /recovery_bin folder...")
		err := os.Mkdir("recovery_bin", 0700)

		if err != nil {
			log.Fatalf("Error creating recovery_bin folder!")
		}
	}
}
