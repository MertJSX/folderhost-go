package utils

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var (
	clients       = make(map[*websocket.Conn]string) // conn -> path
	clientsMu     sync.RWMutex
	connMutexes   = make(map[*websocket.Conn]*sync.Mutex)
	connMutexesMu sync.RWMutex
)

func AddClient(conn *websocket.Conn, path string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[conn] = path

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

	for conn, clientPath := range clients {
		if clientPath == path && conn != exclude {
			go safeWriteMessage(conn, mt, message)
		}
	}
}

func SendToAll(path string, mt int, message []byte) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	var wg sync.WaitGroup
	for conn, clientPath := range clients {
		if clientPath == path {
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
	for _, clientPath := range clients {
		if clientPath == path {
			count++
		}
	}
	return count
}
