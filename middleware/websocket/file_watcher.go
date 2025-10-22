package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/fasthttp/websocket"
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
)

func WatcherDestroyer(path string, ch chan bool) {
	for {
		if utils.GetClientsCount(path) == 0 {
			cache.EditorWatcherCache.Delete(path)
			ch <- true
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func SetupWatcher(path string, stopChan chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Watcher creation error:", err)
		return
	}
	defer watcher.Close()
	defer fmt.Println("Stopped watching...")
	defer close(stopChan)

	err = watcher.Add(path)
	if err != nil {
		log.Println("Watcher add error:", err)
		return
	}

	log.Printf("👀 Watching %s file...\n", path)

	// On file changed by the host computer
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			watcherCache, ok := cache.EditorWatcherCache.Get(path)
			if !ok {
				return
			}
			onFileChanged(event, path, watcherCache)

		case _, ok := <-watcher.Errors:
			if !ok {
				return
			}

		case <-stopChan:
			fmt.Printf("🛑 Watcher received stop signal: %s\n", path)
			return
		}
	}

}

func onFileChanged(event fsnotify.Event, path string, watcherCache types.EditorWatcherCache) {
	if !event.Has(fsnotify.Write) {
		return
	}

	if watcherCache.IsWriting {
		return
	}

	info, err := os.Stat(path)

	if err != nil {
		return
	}

	if info.ModTime() != watcherCache.LastModTime {
		// file changed by host computer
		fmt.Println("File changed by host PC!")
		watcherCache.LastModTime = info.ModTime()
		cache.EditorWatcherCache.SetWithoutTTL(path, watcherCache)
		// Send full content if filesize is lower than 200 KB
		if info.Size() < 200*1024 {
			sendFullContent(path)
		} else {
			sendReloadNotification(path)
		}
	}
}

func sendFullContent(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	msg, _ := json.Marshal(fiber.Map{
		"type":      "full-content-update",
		"content":   string(content),
		"timestamp": time.Now().Unix(),
	})

	utils.SendToAll(path, websocket.TextMessage, msg)
}

func sendReloadNotification(path string) {
	msg, _ := json.Marshal(fiber.Map{
		"type":      "file-changed-externally",
		"timestamp": time.Now().Unix(),
	})

	utils.SendToAll(path, websocket.TextMessage, msg)
}
