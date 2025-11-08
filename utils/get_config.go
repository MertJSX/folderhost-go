package utils

import (
	"log"
	"os"
	"strings"

	"github.com/MertJSX/folder-host-go/utils/config"
	"gopkg.in/yaml.v3"
)

func GetConfig() {
	fileData, err := os.ReadFile("./config.yml")

	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = yaml.Unmarshal(fileData, &config.Config)
	if err != nil {
		log.Fatalf("Config.yml parse error: %v", err)
	}

	config.Config.SizeBytes = ConvertStringToBytes(config.Config.StorageLimit)

	config.Config.Folder = strings.TrimPrefix(config.Config.Folder, "./")
}
