package utils

import (
	"fmt"
	"regexp"

	"github.com/MertJSX/folder-host-go/utils/config"
)

func ReplaceHostPrefix(input string) string {
	pattern := regexp.MustCompile(fmt.Sprintf(`(^|\/)%s(\/|$)`, config.Config.Folder))

	if pattern.MatchString(input) {
		return pattern.ReplaceAllString(input, "${1}./")
	}
	return input
}
