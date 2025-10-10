package types

import (
	"sync"
	"time"
)

type UserDebounceState struct {
	Timer    *time.Timer
	FilePath string
	Username string
	Mu       sync.Mutex
}
