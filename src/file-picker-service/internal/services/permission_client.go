package services

import (
	"context"
	"fmt"
	"time"

	"file-picker-service/proto/generated/permission"
	"google.golang.org/grpc"
)

type PermissionClient struct {
	client permission.PermissionServiceClient
	conn   *grpc.ClientConn // Store connection for proper shutdown
}

// NewPermissionClient creates a new gRPC client for the Permission Service using the recommended grpc.NewClientConn.
func NewPermissionClient(permissionServiceAddress string) (*PermissionClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use NewClientConn to create the connection
	conn, err := grpc.DialContext(ctx, permissionServiceAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to permission service: %v", err)
	}

	client := permission.NewPermissionServiceClient(conn)
	return &PermissionClient{client: client, conn: conn}, nil
}

// Close gracefully closes the gRPC connection
func (p *PermissionClient) Close() error {
	return p.conn.Close()
}

// GrantOwnerPermissions grants full permissions to the owner of the file.
func (p *PermissionClient) GrantOwnerPermissions(ownerID uint64, fileIDs []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &permission.UpdatePermissionRequest{
		OwnerId:     ownerID,
		FileIds:     fileIDs,
		Permissions: []string{"read", "write", "delete"}, // Full permissions for the owner
		IsOwner:     true,                                // Set the owner flag to true
	}

	resp, err := p.client.UpdatePermission(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update permissions for owner: %v", err)
	}

	if !resp.Success {
		return fmt.Errorf("permission service responded with error: %s", resp.Message)
	}

	return nil
}

// ShareFilePermissions grants permissions to a shared user for specified files.
func (p *PermissionClient) ShareFilePermissions(ownerID, sharedUserID uint64, fileIDs []string, permissions []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &permission.UpdatePermissionRequest{
		OwnerId:      ownerID,
		SharedUserId: sharedUserID,
		FileIds:      fileIDs,
		Permissions:  permissions, // Permissions to grant (read, write, etc.)
		IsOwner:      false,        // The file is being shared, not owned by this user
	}

	resp, err := p.client.UpdatePermission(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update permissions for shared user: %v", err)
	}

	if !resp.Success {
		return fmt.Errorf("permission service responded with error: %s", resp.Message)
	}

	return nil
}

func (p *PermissionClient) CheckPermission(userID uint, permissionType string, fileID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &permission.CheckPermissionRequest{
		UserId:         uint64(userID),
		FileId:         fileID,
		Permission: permissionType,
	}

	resp, err := p.client.CheckPermission(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to check permission: %v", err)
	}

	return resp.HasPermission, nil
}