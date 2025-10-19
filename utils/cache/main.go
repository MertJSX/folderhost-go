package cache

import (
	"encoding/json"
	"time"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](5 * time.Minute)
var DirectoryCache *Cache[string, types.ReadDirCache] = CreateCache[string, types.ReadDirCache](30 * time.Second)
var EditorWatcherCache *Cache[string, types.EditorWatcherCache] = CreateCache[string, types.EditorWatcherCache](0)

func ListenDirectorySetCacheEvents() {
	msg, _ := json.Marshal(fiber.Map{
		"type": "directory-update",
	})

	for key := range DirectoryCache.SetCacheEvent {
		utils.SendToAll(key, 1, msg)
	}
}
