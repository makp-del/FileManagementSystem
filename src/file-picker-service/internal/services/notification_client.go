package services

import (
	"context"
	"fmt"

	pb "file-picker-service/proto/generated/notification"
	"google.golang.org/grpc"
)

// NotificationClient defines the gRPC client for Notification Service
type NotificationClient struct {
	client pb.NotificationServiceClient
}

// NewNotificationClient creates a new gRPC client for Notification Service
func NewNotificationClient(conn *grpc.ClientConn) *NotificationClient {
	return &NotificationClient{
		client: pb.NewNotificationServiceClient(conn),
	}
}

// SendNotification sends a notification to the user
func (n *NotificationClient) SendNotification(userID uint, message string) error {
	req := &pb.SendNotificationRequest{
		UserId:  fmt.Sprintf("%d", userID),
		Message: message,
	}

	// Call the Notification Service
	_, err := n.client.SendNotification(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}

	return nil
}