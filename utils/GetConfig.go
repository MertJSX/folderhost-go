package utils

import (
	"log"
	"strings"

	"github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/types"
	"gopkg.in/yaml.v3"
)

var Config types.ConfigFile

func GetConfig() {
	fileData, err := resources.DefaultConfig.ReadFile("default_config.yml")

	if err != nil {
		log.Fatalf("Error reading embedded file: %s", err)
	}

	err = yaml.Unmarshal(fileData, &Config)
	if err != nil {
		log.Fatalf("Config.yml parse error: %v", err)
	}

	Config.Folder = strings.TrimPrefix(Config.Folder, "./")
}
