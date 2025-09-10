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

func SendToAll(path string, mt int, message []byte, exclude *websocket.Conn) {
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
