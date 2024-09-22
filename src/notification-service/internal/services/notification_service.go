package services

import (
	"encoding/json"
	"fmt"
	"notification-service/utils/logger"
	"notification-service/internal/websocket"
)

// NotificationService handles sending notifications to WebSocket clients
type NotificationService struct {
	Hub *websocket.Hub // Reference to the WebSocket hub for managing clients and broadcasting messages
}

// NewNotificationService creates a new instance of NotificationService
func NewNotificationService(hub *websocket.Hub) *NotificationService {
	return &NotificationService{
		Hub: hub,
	}
}

// Notification represents the structure of the notification message
type Notification struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}

// SendNotification sends a notification to all connected WebSocket clients
func (ns *NotificationService) SendNotification(userID uint64, message string) error {
	// Create the notification object
	notification := Notification{
		UserID:  userID,
		Message: message,
	}

	// Convert the notification to JSON
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		logger.Error.Println("Failed to marshal notification to JSON:", err)
		return err
	}

	// Broadcast the notification to all clients in the WebSocket hub
	ns.Hub.Broadcast(notificationJSON)

	logger.Info.Println(fmt.Sprintf("Notification sent to user %d: %s", userID, message))

	return nil
}
