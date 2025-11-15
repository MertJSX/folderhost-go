package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

func FindAvailablePort(startPort, maxPort int) (int, error) {
	for port := startPort; port <= maxPort; port++ {
		if IsPortAvailable(port) {
			log.Printf("Port %d is available!", port)
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available ports found in range %d-%d", startPort, maxPort)
}

func IsPortAvailable(port int) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", fmt.Sprintf("%d", port)), timeout)
	if err != nil {
		return true
	}
	if conn != nil {
		conn.Close()
		return false
	}
	return true
}
