package utils

func LastChar(s string) string {
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	return string(runes[len(runes)-1])
}
