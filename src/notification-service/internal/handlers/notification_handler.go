package handlers

import (
	"context"
	"notification-service/utils/logger"
	"notification-service/internal/services"
	"notification-service/proto/generated/notification"
)

// NotificationHandler implements the gRPC notification service
type NotificationHandler struct {
	notification.UnimplementedNotificationServiceServer // Embed the unimplemented server
	NotificationService *services.NotificationService    // Reference to the notification service
}

// SendNotification handles gRPC requests to send notifications via WebSocket
func (h *NotificationHandler) SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error) {
	logger.Info.Println("Received SendNotification request for user:", req.UserId)

	// Send the notification to the WebSocket clients via NotificationService
	err := h.NotificationService.SendNotification(req.UserId, req.Message)
	if err != nil {
		logger.Error.Println("Failed to send notification:", err)
		return &notification.SendNotificationResponse{
			Success: false,
		}, err
	}

	logger.Info.Println("Notification sent to user:", req.UserId)
	return &notification.SendNotificationResponse{
		Success: true,
	}, nil
}