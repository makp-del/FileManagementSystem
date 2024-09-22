package handlers

import (
	"net/http"
	"notification-service/internal/websocket"
	"notification-service/utils/logger"
)

// WebSocketHandler handles WebSocket connection upgrades and interaction
type WebSocketHandler struct {
	Hub *websocket.Hub // Reference to the hub that manages WebSocket connections
}

// ServeWebSocket upgrades HTTP requests to WebSocket connections
func (h *WebSocketHandler) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("Upgrading HTTP request to WebSocket connection")

	// Upgrade the HTTP request to a WebSocket connection
	websocket.ServeWs(h.Hub, w, r)
}
