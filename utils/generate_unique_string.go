package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUniqueString() string {
	uuid := uuid.New().String()
	cleanUUID := strings.ReplaceAll(uuid, "-", "")
	return cleanUUID[:10]
}
