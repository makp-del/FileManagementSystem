package handlers

import (
	"file-picker-service/internal/models"
	"file-picker-service/internal/services"
	"net/http"
	"file-picker-service/internal/db"

	"github.com/gin-gonic/gin"
)

// FileTransformationHandler handles requests for file transformations (e.g., OCR, image recognition).
func FileTransformationHandler(permissionsClient *services.PermissionClient, transformationClient *services.TransformationClient, notificationClient *services.NotificationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from JWT token
		userID := c.GetUint("userId")

		// Extract file ID from request parameters
		fileID := c.Param("id")

		// Extract transformation type from request parameters
		transformationType := c.Param("type")


		// Check if user has permission to transform the file
		hasPermission, err := permissionsClient.CheckPermission(userID, "transform", fileID)
		if err != nil || !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"msg": "You do not have permission to transform this file"})
			return
		}

		// Get the file path from the database
		filePath, err := models.GetFilePath(db.DB, fileID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to get file path"})
			return
		}

		// Send the transformation request to the transformation service
		//The tranformation type is hardcoded to "ocr" for Optical Character Recognition now, but it could be any other type of transformation
		//such as image recognition, text extraction, etc.
		//The transformation service will process the file and store the result in the same or seperate location.
		//The transformation service will also notify the user once the transformation is complete.
		err = transformationClient.RequestTransformation(fileID, filePath, transformationType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to request file transformation"})
			return
		}

		// Notify the user that the transformation has been initiated
		err = notificationClient.SendNotification(userID, "File transformation has started.")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to send notification"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "File transformation requested successfully"})
	}
}