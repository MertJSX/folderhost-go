package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MertJSX/folder-host-go/utils/config"
)

func ReplaceHostPrefix(input string) string {
	input = strings.ReplaceAll(input, "\\", "/")
	pattern := regexp.MustCompile(fmt.Sprintf(`(^|\/)%s(\/|$)`, config.Config.Folder))

	if pattern.MatchString(input) {
		return pattern.ReplaceAllString(input, "${1}./")
	}
	return input
}
