package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDriveClient struct {
	service *drive.Service
}

// NewGoogleDriveClient creates a new GoogleDriveClient with the provided API key.
func NewGoogleDriveClient(apiKey string) *GoogleDriveClient {
	ctx := context.Background()
	service, err := drive.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create Google Drive service: %v", err)
	}

	return &GoogleDriveClient{
		service: service,
	}
}

// DownloadFile downloads a file from Google Drive and saves it locally.
func (g *GoogleDriveClient) DownloadFile(fileID, authToken string) (string, error) {
	// Create a request to get the file metadata
	file, err := g.service.Files.Get(fileID).Fields("name").Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve file metadata: %v", err)
	}

	// Create a download request for the file
	response, err := g.service.Files.Get(fileID).Download()
	if err != nil {
		return "", fmt.Errorf("unable to download file: %v", err)
	}
	defer response.Body.Close()

	// Save the file locally
	filePath := filepath.Join("./downloads", file.Name)
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to create local file: %v", err)
	}
	defer outFile.Close()

	// Write the downloaded file to the local file system
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		return "", fmt.Errorf("unable to save file locally: %v", err)
	}

	log.Printf("File %s downloaded successfully to %s", file.Name, filePath)
	return filePath, nil
}