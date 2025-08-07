package utils

import (
	"log"
	"strings"

	"github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/types"
	"gopkg.in/yaml.v3"
)

func GetConfig() types.ConfigFile {
	fileData, err := resources.DefaultConfig.ReadFile("default_config.yml")

	if err != nil {
		log.Fatalf("Error reading embedded file: %s", err)
	}

	var config types.ConfigFile

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("Config.yml parse error: %v", err)
	}

	config.Folder = strings.TrimPrefix(config.Folder, "./")

	return config
}
