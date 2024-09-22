package handlers

import (
	"io/ioutil"
	"net/http"

	"file-picker-service/internal/services"

	"github.com/gin-gonic/gin"
)

// FileUploadHandler handles file uploads from the user.
func FileUploadHandler(permissionsClient *services.PermissionClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID and other fields from JWT token (assumed handled by middleware/Istio)
		userID := c.GetUint("userId")

		fileId := c.Param("id")

		// Handle file upload
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Failed to upload file"})
			return
		}

		// Read the file data
		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to read file data"})
			return
		}

		// Convert file data to []byte
		data, err := ioutil.ReadAll(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to read file data"})
			return
		}

		(&services.FileService{}).SaveFile(fileId, userID, file.Filename, data)

		c.JSON(http.StatusOK, gin.H{"msg": "File uploaded successfully"})
	}
}

// ListFilesHandler lists the files that the user has access to.
func ListFilesHandler(permissionsClient *services.PermissionClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from JWT token
		userID := c.GetUint("userId")

		files, error := (&services.FileService{}).ListFiles(userID)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to list files"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"files": files})
	}
}

type FileDownloadRequest struct {
	FileID    string `json:"file_id"`
	Provider  string `json:"provider"`  // e.g., "google_drive", "dropbox"
	AuthToken string `json:"auth_token"`
}

// FileDownloadHandler handles requests to download a file from cloud storage.
func FileDownloadHandler(permissionsClient *services.PermissionClient, downloaderClient *services.DownloaderClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FileDownloadRequest

		userID := c.GetUint("userId")

		// Bind request JSON to struct
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Call the service to download the file
		filePath, err := (&services.FileService{}).DownloadFile(userID, req.FileID, req.Provider, req.AuthToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with the local file path where the file was saved
		c.JSON(http.StatusOK, gin.H{"file_path": filePath})
	}
}

// AddPermissionsHandler allows file owners to share files with another user.
func AddPermissionsHandler(permissionsClient *services.PermissionClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			SharedUserID uint64   `json:"shared_user_id"` // ID of the user to share the file with
			FileIDs      []string `json:"file_ids"`       // List of file IDs to share
			Permissions  []string `json:"permissions"`    // List of permissions (read, write, delete)
		}

		// Parse the request body
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Extract owner ID from JWT (assuming user_id is available via middleware/Istio)
		ownerID := c.GetUint("userId")

		// Use the FileService to add permissions for the shared user
		if err := (&services.FileService{}).ShareFilePermissions(uint64(ownerID), req.SharedUserID, req.FileIDs, req.Permissions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Permissions updated successfully"})
	}
}