package types

import (
	"time"
)

type EditorWatcherCache struct {
	LastModTime time.Time
	IsWriting   bool
}
