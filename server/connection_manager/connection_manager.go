package connectionmanager

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[*websocket.Conn]bool
	mutex       sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (cm *ConnectionManager) AddConnection(conn *websocket.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.connections[conn] = true
}

func (cm *ConnectionManager) RemoveConnection(conn *websocket.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	delete(cm.connections, conn)
}

func (cm *ConnectionManager) Broadcast(message []byte) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for conn := range cm.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("Closing connection with %s", conn.RemoteAddr())
			conn.Close()
			delete(cm.connections, conn)
		}
	}
}
