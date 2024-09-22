package websocket

// Hub maintains the set of active WebSocket connections and broadcasts messages between them
type Hub struct {
	// Registered connections
	clients map[*Connection]bool

	// Inbound messages from the connections
	broadcast chan []byte

	// Register requests from the connections
	register chan *Connection

	// Unregister requests from connections
	unregister chan *Connection
}

// NewHub creates a new Hub instance
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Connection]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
	}
}

// Run runs the hub, handling connection registration, unregistration, and message broadcasting
func (h *Hub) Run() {
	for {
		select {
		case connection := <-h.register:
			h.clients[connection] = true
		case connection := <-h.unregister:
			if _, ok := h.clients[connection]; ok {
				delete(h.clients, connection)
				close(connection.send)
			}
		case message := <-h.broadcast:
			for connection := range h.clients {
				select {
				case connection.send <- message:
				default:
					close(connection.send)
					delete(h.clients, connection)
				}
			}
		}
	}
}

// Broadcast sends a message to all clients in the hub
func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}