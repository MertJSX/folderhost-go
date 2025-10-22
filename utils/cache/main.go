package cache

import (
	"encoding/json"
	"os"
	"time"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

var SessionCache *Cache[string, types.Account] = CreateCache[string, types.Account](5 * time.Minute)
var DirectoryCache *Cache[string, types.ReadDirCache] = CreateCache[string, types.ReadDirCache](30 * time.Second)
var EditorWatcherCache *Cache[string, types.EditorWatcherCache] = CreateCache[string, types.EditorWatcherCache](0)
var FileContentCache *Cache[string, string] = CreateCache[string, string](500 * time.Millisecond)

func ListenDirectorySetCacheEvents() {
	msg, _ := json.Marshal(fiber.Map{
		"type": "directory-update",
	})

	for key := range DirectoryCache.SetCacheEvent {
		utils.SendToAll(key, 1, msg)
	}
}

func ListenFileContentCacheEvents() {
	for event := range FileContentCache.TimeoutCacheEvent {
		watcherCache, ok := EditorWatcherCache.Get(event.Key)
		if !ok {
			return
		}

		watcherCache.IsWriting = true

		EditorWatcherCache.SetWithoutTTL(event.Key, watcherCache)

		err := os.WriteFile(event.Key, []byte(event.Data), 0644)

		if err != nil {
			return
		}

		fileStat, err := os.Stat(event.Key)

		if err != nil {
			return
		}

		watcherCache.LastModTime = fileStat.ModTime()
		watcherCache.IsWriting = false

		EditorWatcherCache.SetWithoutTTL(event.Key, watcherCache)

		if directoryCache, ok := DirectoryCache.Get(utils.GetParentPath(event.Key) + "/"); ok {
			for index, v := range directoryCache.Items {
				v.SizeBytes = fileStat.Size()
				directoryCache.Items[index] = v
			}
			DirectoryCache.Set(utils.GetParentPath(event.Key)+"/", directoryCache, 600)
		}

	}
}
