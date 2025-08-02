package utils

import (
	"fmt"
	"math"
)

func ConvertBytesToString(bytes int64) string {
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}
	var bytesFloat float64 = float64(bytes)

	if bytes == 0 {
		return "N/A"
	}

	var i int = int(math.Floor(math.Log(bytesFloat) / math.Log(1024)))

	if i == 0 {
		return fmt.Sprintf("%d %s", bytes, sizes[i])
	}

	return fmt.Sprintf("%.1f %s", (bytesFloat / math.Pow(1024, float64(i))), sizes[i])
}
