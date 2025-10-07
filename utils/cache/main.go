package cache

import (
	"time"

	"github.com/MertJSX/folder-host-go/types"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](*time.NewTicker(1 * time.Minute))
