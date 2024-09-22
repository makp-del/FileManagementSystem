package websocket

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Constants for connection settings
const (
	writeWait = 10 * time.Second        // Time allowed to write a message to the client
	pongWait  = 60 * time.Second        // Time allowed to read the next pong message from the client
	pingPeriod = (pongWait * 9) / 10    // Send pings to client at this period. Must be less than pongWait
	maxMessageSize = 512                // Maximum message size allowed from the client
)

// upgrader upgrades HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Bypasses CORS for WebSocket connections
}

// Connection represents a single WebSocket connection
type Connection struct {
	hub  *Hub          // The hub managing all WebSocket connections
	conn *websocket.Conn // The actual WebSocket connection
	send chan []byte     // Buffered channel of outbound messages
}

// readPump pumps messages from the WebSocket connection to the hub
// It runs in a goroutine for each connection
func (c *Connection) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the WebSocket connection
// It runs in a goroutine for each connection
func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued messages to the current WebSocket message
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles WebSocket requests from clients
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	connection := &Connection{hub: hub, conn: conn, send: make(chan []byte, 256)}
	connection.hub.register <- connection

	// Start goroutines to handle WebSocket connection
	go connection.writePump()
	go connection.readPump()
}