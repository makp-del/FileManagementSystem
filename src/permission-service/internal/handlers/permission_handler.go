package handlers

import (
	"context"
	"permission-service/utils/logger"
	"permission-service/internal/services"
	"permission-service/proto/generated/permission"
)

// PermissionHandler implements the PermissionServiceServer
type PermissionHandler struct {
	permission.UnimplementedPermissionServiceServer // Embed the unimplemented server
	PermissionService *services.PermissionService
}

// CheckPermission checks if a user has a specific permission for a file.
func (h *PermissionHandler) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (*permission.CheckPermissionResponse, error) {
	logger.Info.Println("Received CheckPermission request for user:", req.UserId, "on file:", req.FileId)

	// Delegate the logic to the service
	hasPermission, err := h.PermissionService.CheckPermission(req.UserId, req.FileId, req.Permission)
	if err != nil {
		return nil, err
	}

	return &permission.CheckPermissionResponse{
		HasPermission: hasPermission,
	}, nil
}

// UpdatePermission updates permissions for a file or folder for a shared user.
func (h *PermissionHandler) UpdatePermission(ctx context.Context, req *permission.UpdatePermissionRequest) (*permission.UpdatePermissionResponse, error) {
	logger.Info.Println("Received UpdatePermission request for owner:", req.OwnerId, "to share with user:", req.SharedUserId)

	// Delegate the logic to the service
	err := h.PermissionService.UpdatePermission(req.OwnerId, req.SharedUserId, req.FileIds, req.Permissions)
	if err != nil {
		return &permission.UpdatePermissionResponse{
			Success: false,
			Message: "Failed to update permissions",
		}, err
	}

	return &permission.UpdatePermissionResponse{
		Success: true,
		Message: "Permissions updated successfully",
	}, nil
}