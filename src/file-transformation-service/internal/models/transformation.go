package models

// TransformationType represents the different types of transformations available
type TransformationType string

const (
	ResizeTransformation  TransformationType = "resize"
	ConvertTransformation TransformationType = "convert"
)

// TransformationRequest represents a file transformation request
type TransformationRequest struct {
	FilePath           string            // Path of the file to be transformed
	TransformationType TransformationType // Type of transformation (e.g., resize, convert)
	Options            map[string]string  // Additional options for transformation (e.g., dimensions for resizing)
}

// TransformationResponse represents the response after a transformation is completed
type TransformationResponse struct {
	FilePath string // Path of the transformed file (should be the same as the input file path)
	Status   string // Status of the transformation (success, failure)
	Message  string // Optional message (e.g., error details or success message)
}

// NewTransformationRequest creates a new TransformationRequest with provided parameters
func NewTransformationRequest(filePath string, transformationType TransformationType, options map[string]string) *TransformationRequest {
	return &TransformationRequest{
		FilePath:           filePath,
		TransformationType: transformationType,
		Options:            options,
	}
}

// NewTransformationResponse creates a new TransformationResponse with provided parameters
func NewTransformationResponse(filePath string, status string, message string) *TransformationResponse {
	return &TransformationResponse{
		FilePath: filePath,
		Status:   status,
		Message:  message,
	}
}

// IsValid checks if the transformation request is valid
func (req *TransformationRequest) IsValid() bool {
	if req.FilePath == "" {
		return false
	}
	if req.TransformationType != ResizeTransformation && req.TransformationType != ConvertTransformation {
		return false
	}
	return true
}