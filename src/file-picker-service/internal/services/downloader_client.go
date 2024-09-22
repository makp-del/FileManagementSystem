package services

import (
	"context"
	"fmt"
	"time"

	"file-picker-service/proto/generated/filedownloader"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DownloaderClient struct {
	client filedownloader.FileDownloaderServiceClient
}

func NewDownloaderClient(downloaderServiceAddress string) (*DownloaderClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, downloaderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to file-downloader-service: %v", err)
	}

	client := filedownloader.NewFileDownloaderServiceClient(conn)
	return &DownloaderClient{client: client}, nil
}

func (d *DownloaderClient) DownloadFile(fileID, provider, authToken string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &filedownloader.DownloadFileRequest{
		FileId:    fileID,
		Provider:  provider,
		AuthToken: authToken,
	}

	resp, err := d.client.DownloadFile(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %v", err)
	}

	return resp.FilePath, nil
}