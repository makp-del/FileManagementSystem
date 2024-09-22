package services

import (
	"context"
	"fmt"

	pb "file-picker-service/proto/generated/transformation"
	"google.golang.org/grpc"
)

// TransformationClient defines the gRPC client for File Transformation Service
type TransformationClient struct {
	client pb.FileTransformationServiceClient
}

// NewTransformationClient creates a new gRPC client for File Transformation Service
func NewTransformationClient(conn *grpc.ClientConn) *TransformationClient {
	return &TransformationClient{
		client: pb.NewFileTransformationServiceClient(conn),
	}
}

// RequestTransformation sends a request to transform the file
func (t *TransformationClient) RequestTransformation(fileID string, filePath string, tranformationType string) error {
	req := &pb.TransformFileRequest{
		FileId: fileID,
		FilePath: filePath,
		TransformationType: tranformationType,
	}

	// Call the File Transformation Service
	_, err := t.client.TransformFile(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to request file transformation: %v", err)
	}

	return nil
}