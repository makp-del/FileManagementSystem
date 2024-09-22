package services

import (
	"errors"
	"file-picker-service/internal/db"
	"file-picker-service/internal/models"
	"os"
	"path/filepath"
)

// FileService provides core business logic for file operations.
type FileService struct {
	permissionsClient    PermissionClient
	downloaderClient     DownloaderClient
	transformationClient TransformationClient
	notificationClient   NotificationClient
}

// NewFileService creates a new FileService instance with required gRPC clients.
func NewFileService(permissionsClient PermissionClient, downloaderClient DownloaderClient, transformationClient TransformationClient, notificationClient NotificationClient) *FileService {
	return &FileService{
		permissionsClient:    permissionsClient,
		downloaderClient:     downloaderClient,
		transformationClient: transformationClient,
		notificationClient:   notificationClient,
	}
}

// SaveFile handles saving the uploaded file and metadata to the database and storage.
func (s *FileService) SaveFile(fileId string, userID uint, fileName string, fileData []byte) error {

	// Check write permissions via Permissions Service
	hasPermission, err := s.permissionsClient.CheckPermission(userID, "write", "")
	if err != nil || !hasPermission {
		return errors.New("user does not have write permission")
	}

	//Give owner permission to the user who uploaded the file
	err = s.permissionsClient.GrantOwnerPermissions(uint64(userID), []string{fileId})
	if err != nil {
		return err
	}

	// Save file metadata to the database
	err = models.SaveFileMetadata(db.DB, fileId, fileName, userID)
	if err != nil {
		return err
	}

	// Save the file to local storage (or cloud storage if configured)
	filePath := filepath.Join("./uploads", fileName)
	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return err
	}

	// Notify the user that the file has been saved
	err = s.notificationClient.SendNotification(userID, "Your file has been saved.")
	if err != nil {
		return err
	}

	return nil
}

// ListFiles retrieves a list of files that the user has access to (owned or shared).
func (s *FileService) ListFiles(userID uint) ([]models.File, error) {
	// Check read permissions via Permissions Service
	hasPermission, err := s.permissionsClient.CheckPermission(userID, "read", "")
	if err != nil || !hasPermission {
		return nil, errors.New("user does not have read permission")
	}

	// Fetch the list of files from the database
	files, err := models.ListFilesByUser(db.DB, userID)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// DownloadFile retrieves the file using the File-Downloader-Service.
func (s *FileService) DownloadFile(userID uint, fileID, provider, authToken string) (string, error) {
	// Check read permission for the file
	hasPermission, err := s.permissionsClient.CheckPermission(userID, "write", fileID)
	if err != nil || !hasPermission {
		return "", errors.New("user does not have permission to upload this file")
	}

	// Use the File-Downloader-Service to get the file path
	filePath, err := s.downloaderClient.DownloadFile(fileID, provider, authToken)
	if err != nil {
		return "", err
	}

	//Notify the user that the download has started
	err = s.notificationClient.SendNotification(userID, "Your file download has started.")
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// TransformFile requests a file transformation and notifies the user.
func (s *FileService) TransformFile(userID uint, fileID string, filePath string, transformationType string) error {
	// Check permission to transform the file
	hasPermission, err := s.permissionsClient.CheckPermission(userID, "transform", fileID)
	if err != nil || !hasPermission {
		return errors.New("user does not have permission to transform this file")
	}

	// Request file transformation via Transformation Service
	err = s.transformationClient.RequestTransformation(fileID, filePath, transformationType)
	if err != nil {
		return err
	}

	// Notify the user that the transformation process has started
	err = s.notificationClient.SendNotification(userID, "Your file transformation has started.")
	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) ShareFilePermissions(ownerID, sharedUserID uint64, fileIDs []string, permissions []string) error {
	// Use the Permission Client to update permissions for the shared user
	err := s.permissionsClient.ShareFilePermissions(ownerID, sharedUserID, fileIDs, permissions)
	if err != nil {
		return errors.New("failed to update permissions for shared user: %v")
	}

	return nil
}