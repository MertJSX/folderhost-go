package websocket

import (
	"encoding/json"
	"log"
	"net/url"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func HandleWebsocket(c *websocket.Conn) {
	var path string = c.Params("path")

	path, err := url.PathUnescape(path)
	if err != nil {
		c.Close()
		return
	}

	config := &config.Config
	path = config.Folder + path

	if !utils.IsSafePath(path) {
		c.Close()
		return
	}

	fileStat, err := os.Stat(path)
	if err != nil {
		return
	}

	if utils.IsExistingWSConnectionPath(path) || fileStat.IsDir() {
		utils.AddClient(c, path, fileStat.IsDir())
	} else {
		freeSpace, _ := utils.GetRemainingFolderSpace()
		if freeSpace >= (200 * 1024) {
			utils.AddClient(c, path, fileStat.IsDir())
		} else {
			return
		}
	}

	updateClientsCount(path)

	defer updateClientsCount(path)
	defer utils.RemoveClient(c)
	defer log.Printf("WebSocket disconnected - User: %s, Path: %s\n", c.Locals("username").(string), path)
	defer c.Close()

	var username string = c.Locals("username").(string)
	defer utils.TriggerPendingLog(username, path)

	_, ok := cache.EditorWatcherCache.Get(path)

	if !ok {
		cache.EditorWatcherCache.SetWithoutTTL(path, types.EditorWatcherCache{
			LastModTime: fileStat.ModTime(),
			IsWriting:   false,
		})
		if !fileStat.IsDir() {
			channel := make(chan bool, 1)
			go SetupWatcher(path, channel)
			defer func() {
				go WatcherDestroyer(path, channel)
			}()
		}
	}

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			return
		}

		if err := processWebSocketMessage(msg, path, c, mt); err != nil {
			log.Println("Message processing error:", err)
		}
	}
}

func updateClientsCount(path string) {
	clientsCount, err := json.Marshal(fiber.Map{
		"type":  "editor-update-usercount",
		"count": utils.GetClientsCount(path),
	})

	if err == nil {
		utils.SendToAll(path, 1, clientsCount)
	}
}
