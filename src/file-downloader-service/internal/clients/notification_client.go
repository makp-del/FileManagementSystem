package clients

import (
	"context"
	"fmt"
	"time"

	"file-downloader-service/proto/generated/notification"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	client notification.NotificationServiceClient
}

// NewNotificationClient creates a new gRPC client for the Notification Service.
func NewNotificationClient(notificationServiceAddress string) (*NotificationClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, notificationServiceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to notification service: %v", err)
	}

	client := notification.NewNotificationServiceClient(conn)
	return &NotificationClient{client: client}, nil
}

// SendNotification sends a notification after a file download is completed.
func (n *NotificationClient) SendNotification(userID, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &notification.SendNotificationRequest{
		UserId:  userID,
		Message: message,
	}

	_, err := n.client.SendNotification(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}

	return nil
}