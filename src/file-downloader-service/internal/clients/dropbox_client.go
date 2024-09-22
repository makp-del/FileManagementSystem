package clients

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

type DropboxClient struct {
	client files.Client
}

// NewDropboxClient creates a new DropboxClient with the provided access token.
func NewDropboxClient(accessToken string) *DropboxClient {
	config := dropbox.Config{
		Token:    accessToken,
		LogLevel: dropbox.LogInfo, // Change this to dropbox.LogOff to disable logging
	}
	client := files.New(config)

	return &DropboxClient{
		client: client,
	}
}

// DownloadFile downloads a file from Dropbox and saves it locally.
func (d *DropboxClient) DownloadFile(fileID, authToken string) (string, error) {
	// Make a request to download the file from Dropbox
	arg := files.NewDownloadArg(fileID)
	response, content, err := d.client.Download(arg)
	if err != nil {
		return "", fmt.Errorf("failed to download file from Dropbox: %v", err)
	}
	defer content.Close()

	// Save the file locally
	filePath := filepath.Join("./downloads", response.Name)
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to create local file: %v", err)
	}
	defer outFile.Close()

	// Write the downloaded file to the local file system
	_, err = io.Copy(outFile, content)
	if err != nil {
		return "", fmt.Errorf("unable to save file locally: %v", err)
	}

	log.Printf("File %s downloaded successfully to %s", response.Name, filePath)
	return filePath, nil
}