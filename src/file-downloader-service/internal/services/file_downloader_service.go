package services

import (
	"fmt"
	"log"

	"file-downloader-service/internal/clients"
)

// FileDownloaderService handles the logic for downloading files from cloud storage.
type FileDownloaderService struct {
	googleDriveClient  *clients.GoogleDriveClient
	dropboxClient      *clients.DropboxClient
	notificationClient *clients.NotificationClient
}

// NewFileDownloaderService creates a new instance of FileDownloaderService.
func NewFileDownloaderService(googleDriveClient *clients.GoogleDriveClient, dropboxClient *clients.DropboxClient, notificationClient *clients.NotificationClient) *FileDownloaderService {
	return &FileDownloaderService{
		googleDriveClient:  googleDriveClient,
		dropboxClient:      dropboxClient,
		notificationClient: notificationClient,
	}
}

// DownloadFile downloads a file from the specified cloud provider and saves it locally.
func (s *FileDownloaderService) DownloadFile(fileID, provider, authToken string) (string, error) {
	var filePath string
	var err error

	// Download the file based on the provider (Google Drive or Dropbox)
	switch provider {
	case "google_drive":
		filePath, err = s.googleDriveClient.DownloadFile(fileID, authToken)
	case "dropbox":
		filePath, err = s.dropboxClient.DownloadFile(fileID, authToken)
	default:
		return "", fmt.Errorf("unsupported provider: %s", provider)
	}

	if err != nil {
		log.Printf("Error downloading file: %v", err)
		return "", err
	}

	// Notify the user that the file has been downloaded
	err = s.notificationClient.SendNotification("user_id_placeholder", fmt.Sprintf("File %s has been downloaded to %s", fileID, filePath))
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
		return "", err
	}

	log.Printf("File %s downloaded successfully to %s", fileID, filePath)
	return filePath, nil
}