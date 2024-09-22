package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"file-transformation-service/internal/utils"
	"file-transformation-service/proto/generated/file_transformation"
)

// FileTransformationHandler implements the gRPC service for file transformation
type FileTransformationHandler struct {
	file_transformation.UnimplementedFileTransformationServiceServer
	Timeout int // Timeout for transformation operations in seconds
}

// TransformFile handles the file transformation request
func (h *FileTransformationHandler) TransformFile(ctx context.Context, req *file_transformation.TransformFileRequest) (*file_transformation.TransformFileResponse, error) {
	pkg.Logger.Infof("Received request to transform file: %s with transformation type: %s", req.FilePath, req.TransformationType)

	// Check if the file exists
	if _, err := os.Stat(req.FilePath); errors.Is(err, os.ErrNotExist) {
		msg := fmt.Sprintf("File not found: %s", req.FilePath)
		pkg.Logger.Error(msg)
		return &file_transformation.TransformFileResponse{
			FilePath: req.FilePath,
			Status:   "failure",
			Message:  msg,
		}, nil
	}

	// Handle transformation based on type (e.g., resize, format conversion)
	switch req.TransformationType {
	case "resize":
		// Example: Perform image resizing
		if err := h.resizeFile(req.FilePath, req.Options); err != nil {
			pkg.Logger.Errorf("Failed to resize file: %v", err)
			return &file_transformation.TransformFileResponse{
				FilePath: req.FilePath,
				Status:   "failure",
				Message:  fmt.Sprintf("Failed to resize file: %v", err),
			}, nil
		}
		pkg.Logger.Infof("File resized successfully: %s", req.FilePath)

	case "convert":
		// Example: Perform file format conversion
		if err := h.convertFileFormat(req.FilePath, req.Options); err != nil {
			pkg.Logger.Errorf("Failed to convert file format: %v", err)
			return &file_transformation.TransformFileResponse{
				FilePath: req.FilePath,
				Status:   "failure",
				Message:  fmt.Sprintf("Failed to convert file format: %v", err),
			}, nil
		}
		pkg.Logger.Infof("File format converted successfully: %s", req.FilePath)

	default:
		// Unsupported transformation type
		msg := fmt.Sprintf("Unsupported transformation type: %s", req.TransformationType)
		pkg.Logger.Error(msg)
		return &file_transformation.TransformFileResponse{
			FilePath: req.FilePath,
			Status:   "failure",
			Message:  msg,
		}, nil
	}

	// Simulate transformation delay (optional)
	time.Sleep(time.Duration(h.Timeout) * time.Second)

	// Return success response
	return &file_transformation.TransformFileResponse{
		FilePath: req.FilePath,
		Status:   "success",
		Message:  "Transformation completed successfully",
	}, nil
}

// resizeFile handles resizing of an image file
func (h *FileTransformationHandler) resizeFile(filePath string, options map[string]string) error {
	width := options["width"]
	height := options["height"]

	// Simulate resizing logic (replace with actual resizing code)
	pkg.Logger.Infof("Resizing file: %s to width: %s, height: %s", filePath, width, height)

	// Example: Replace with actual image resizing code using a library like "github.com/nfnt/resize"
	return nil
}

// convertFileFormat handles file format conversion (e.g., from .png to .jpg)
func (h *FileTransformationHandler) convertFileFormat(filePath string, options map[string]string) error {
	newFormat := options["format"]

	// Simulate format conversion logic (replace with actual conversion code)
	pkg.Logger.Infof("Converting file: %s to format: %s", filePath, newFormat)

	// Example: Replace with actual file format conversion code
	return nil
}