package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/MertJSX/folder-host-go/types"
	"gopkg.in/yaml.v3"
)

func GetConfig() types.ConfigFile {
	fileData, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}
	var config types.ConfigFile

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("Config.yml parse error: %v", err)
	}

	r, _ := regexp.Compile(`^\./`)

	if r.MatchString(config.Folder) {
		config.Folder = config.Folder[2:]
	}

	return config
}
