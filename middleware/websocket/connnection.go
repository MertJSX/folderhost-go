package websocket

import (
	"encoding/json"
	"log"
	"net/url"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
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

	config := &utils.Config
	path = config.Folder + path

	if !utils.IsSafePath(path) {
		c.Close()
		return
	}

	utils.AddClient(c, path)
	updateClientsCount(path)

	defer updateClientsCount(path)
	defer utils.RemoveClient(c)
	defer log.Printf("WebSocket disconnected - User: %s, Path: %s\n", c.Locals("username").(string), path)
	defer c.Close()

	var username string = c.Locals("username").(string)
	defer utils.TriggerPendingLog(username, path)

	_, ok := cache.EditorWatcherCache.Get(path)
	fileStat, err := os.Stat(path)
	if err != nil {
		return
	}

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
