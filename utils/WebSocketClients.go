package utils

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var (
	clients   = make(map[*websocket.Conn]string) // conn -> path
	clientsMu sync.RWMutex
)

func AddClient(conn *websocket.Conn, path string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[conn] = path
}

func RemoveClient(conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, conn)
}

func SendToAllExclude(path string, mt int, message []byte, exclude *websocket.Conn) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	fmt.Println("Send to all operation")
	for conn, clientPath := range clients {
		fmt.Printf("Client: %s\n", clientPath)
		if clientPath == path && conn != exclude {
			fmt.Println("Msg was sent!")
			conn.WriteMessage(mt, message)
		}
	}
}

func SendToAll(path string, mt int, message []byte) {
	clientsMu.RLock()
	defer clientsMu.RUnlock()

	fmt.Println("Send to all operation")
	for conn, clientPath := range clients {
		fmt.Printf("Client: %s\n", clientPath)
		if clientPath == path {
			fmt.Println("Msg was sent!")
			conn.WriteMessage(mt, message)
		}
	}
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
