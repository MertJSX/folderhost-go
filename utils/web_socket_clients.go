package utils

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var (
	clients       = make(map[*websocket.Conn]ClientInfo) // conn -> path
	clientsMu     sync.RWMutex
	connMutexes   = make(map[*websocket.Conn]*sync.Mutex)
	connMutexesMu sync.RWMutex
)

type ClientInfo struct {
	Path        string
	IsDirectory bool
}

func AddClient(conn *websocket.Conn, path string, isDirectory bool) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[conn] = ClientInfo{
		Path:        path,
		IsDirectory: isDirectory,
	}

	connMutexesMu.Lock()
	defer connMutexesMu.Unlock()
	connMutexes[conn] = &sync.Mutex{}
}

func RemoveClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, conn)

	connMutexesMu.Lock()
	defer connMutexesMu.Unlock()
	delete(connMutexes, conn)
}

func safeWriteMessage(conn *websocket.Conn, mt int, message []byte) error {
	connMutexesMu.RLock()
	mutex, exists := connMutexes[conn]
	connMutexesMu.RUnlock()

	if !exists {
		return nil
	}

	mutex.Lock()
	defer mutex.Unlock()
	return conn.WriteMessage(mt, message)
}

func SendToAllExclude(path string, mt int, message []byte, exclude *websocket.Conn) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	for conn, client := range clients {
		if client.Path == path && conn != exclude {
			go safeWriteMessage(conn, mt, message)
		}
	}
}

func SendToAll(path string, mt int, message []byte) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	var wg sync.WaitGroup
	for conn, client := range clients {
		if client.Path == path {
			wg.Add(1)
			go func(c *websocket.Conn) {
				defer wg.Done()
				safeWriteMessage(c, mt, message)
			}(conn)
		}
	}
	wg.Wait()
}

func GetClientsCount(path string) int {
	clientsMu.RLock()
	defer clientsMu.RUnlock()
	count := 0
	for _, client := range clients {
		if client.Path == path {
			count++
		}
	}
	return count
}

func GetActiveFileCount() int {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	uniqueFiles := make(map[string]bool)
	for _, clientInfo := range clients {
		if !clientInfo.IsDirectory {
			uniqueFiles[clientInfo.Path] = true
		}
	}
	return len(uniqueFiles)
}

func IsExistingWSConnectionPath(path string) bool {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	for _, clientInfo := range clients {
		if clientInfo.Path == path {
			return true
		}
	}

	return false
}
