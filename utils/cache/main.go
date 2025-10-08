package cache

import (
	"time"

	"github.com/MertJSX/folder-host-go/types"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](*time.NewTicker(5 * time.Minute))
var DirectoryCache *Cache[string, types.ReadDirCache] = CreateCache[string, types.ReadDirCache](*time.NewTicker(30 * time.Second))
