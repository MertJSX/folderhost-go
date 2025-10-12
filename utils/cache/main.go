package cache

import (
	"time"

	"github.com/MertJSX/folder-host-go/types"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](5 * time.Minute)
var DirectoryCache *Cache[string, types.ReadDirCache] = CreateCache[string, types.ReadDirCache](30 * time.Second)
var EditorWatcherCache *Cache[string, types.EditorWatcherCache] = CreateCache[string, types.EditorWatcherCache](0)
var RemainingFolderSpace int64 = 0
