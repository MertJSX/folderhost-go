package cache

import (
	"encoding/json"
	"time"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](5*time.Minute, CacheProperties{
	SetCacheEvent:     false,
	TimeoutCacheEvent: false,
})
var DirectoryCache *Cache[string, types.ReadDirCache] = CreateCache[string, types.ReadDirCache](30*time.Second, CacheProperties{
	SetCacheEvent:     true,
	TimeoutCacheEvent: false,
})

var EditorWatcherCache *Cache[string, types.EditorWatcherCache] = CreateCache[string, types.EditorWatcherCache](0, CacheProperties{
	SetCacheEvent:     false,
	TimeoutCacheEvent: false,
})

var DownloadLinkCache *Cache[string, types.DownloadLinkCache] = CreateCache[string, types.DownloadLinkCache](1*time.Minute, CacheProperties{
	SetCacheEvent:     false,
	TimeoutCacheEvent: false,
})

func ListenDirectorySetCacheEvents() {
	msg, _ := json.Marshal(fiber.Map{
		"type": "directory-update",
	})

	for key := range DirectoryCache.SetCacheEvent {
		utils.SendToAll(key, 1, msg)
	}
}
