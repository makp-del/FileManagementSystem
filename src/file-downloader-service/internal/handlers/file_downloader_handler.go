package handlers

import (
	"context"
	"log"

	"file-downloader-service/internal/services"
	"file-downloader-service/proto/generated/filedownloader"
)

// FileDownloaderHandler is the gRPC handler for the file-downloader-service.
type FileDownloaderHandler struct {
	filedownloader.UnimplementedFileDownloaderServiceServer
	downloaderService *services.FileDownloaderService
}

// NewFileDownloaderHandler creates a new FileDownloaderHandler.
func NewFileDownloaderHandler(downloaderService *services.FileDownloaderService) *FileDownloaderHandler {
	return &FileDownloaderHandler{
		downloaderService: downloaderService,
	}
}

// DownloadFile handles the gRPC request for downloading a file from cloud storage.
func (h *FileDownloaderHandler) DownloadFile(ctx context.Context, req *filedownloader.DownloadFileRequest) (*filedownloader.DownloadFileResponse, error) {
	log.Printf("Received request to download file: %s from provider: %s", req.FileId, req.Provider)

	// Call the service to download the file
	filePath, err := h.downloaderService.DownloadFile(req.FileId, req.Provider, req.AuthToken)
	if err != nil {
		log.Printf("Failed to download file: %v", err)
		return nil, err
	}

	// Return the file path in the response
	return &filedownloader.DownloadFileResponse{
		FilePath: filePath,
	}, nil
}