package utils

import (
	"fmt"
	"regexp"
)

func ReplaceHostPrefix(input string) string {
	pattern := regexp.MustCompile(fmt.Sprintf(`(^|\/)%s(\/|$)`, Config.Folder))

	if pattern.MatchString(input) {
		return pattern.ReplaceAllString(input, "${1}./")
	}
	return input
}
