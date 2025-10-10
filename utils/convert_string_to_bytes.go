package utils

import (
	"math"
	"strconv"
	"strings"
)

// Missing possible error handlings... But still should work...
func ConvertStringToBytes(size string) int64 {
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}

	parts := strings.Split(size, " ")
	value, err := strconv.ParseFloat(parts[0], 64)

	if err != nil {
		return 0
	}

	var unit string = parts[1]
	var index int = IndexOf(unit, sizes)

	if index == -1 {
		return 0
	}

	return int64(value * math.Pow(1024, float64(index)))

}
