package utils

import (
	"fmt"
	"regexp"

	"github.com/MertJSX/folder-host-go/types"
)

func ReplaceHostPrefix(input string, config types.ConfigFile) string {
	pattern := regexp.MustCompile(fmt.Sprintf(`(^|\/)%s(\/|$)`, config.Folder))

	if pattern.MatchString(input) {
		return pattern.ReplaceAllString(input, "${1}./")
	}
	return input
}
