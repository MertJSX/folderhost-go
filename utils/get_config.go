package utils

import (
	"log"
	"os"
	"strings"

	"github.com/MertJSX/folder-host-go/types"
	"gopkg.in/yaml.v3"
)

var Config types.ConfigFile

func GetConfig() {
	fileData, err := os.ReadFile("./config.yml")

	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = yaml.Unmarshal(fileData, &Config)
	if err != nil {
		log.Fatalf("Config.yml parse error: %v", err)
	}

	Config.SizeBytes = ConvertStringToBytes(Config.StorageLimit)

	Config.Folder = strings.TrimPrefix(Config.Folder, "./")
}
