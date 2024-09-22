package services

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"

	"file-transformation-service/internal/utils"
	"github.com/nfnt/resize"
)

// FileTransformationService defines the service for transforming files
type FileTransformationService struct{}

// NewFileTransformationService creates a new instance of the service
func NewFileTransformationService() *FileTransformationService {
	return &FileTransformationService{}
}

// TransformFile performs the actual file transformation based on the provided type and options
func (s *FileTransformationService) TransformFile(filePath, transformationType string, options map[string]string) error {
	switch transformationType {
	case "resize":
		return s.resizeFile(filePath, options)
	case "convert":
		return s.convertFileFormat(filePath, options)
	default:
		return fmt.Errorf("unsupported transformation type: %s", transformationType)
	}
}

// resizeFile performs the file resizing based on the provided options
func (s *FileTransformationService) resizeFile(filePath string, options map[string]string) error {
	// Ensure required options (e.g., width and height) are provided
	widthStr, ok := options["width"]
	if !ok {
		return fmt.Errorf("missing required option: width")
	}
	heightStr, ok := options["height"]
	if !ok {
		return fmt.Errorf("missing required option: height")
	}

	// Convert width and height to integers
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return fmt.Errorf("invalid width value: %v", err)
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return fmt.Errorf("invalid height value: %v", err)
	}

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the image (supports PNG and JPEG)
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}
	pkg.Logger.Infof("Resizing image of format: %s", format)

	// Perform the resize operation using Lanczos3 algorithm
	newImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// Create a new file to store the resized image (overwrite the original)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Encode and save the resized image based on the original format
	if format == "jpeg" {
		err = jpeg.Encode(out, newImg, &jpeg.Options{Quality: 85})
	} else if format == "png" {
		err = png.Encode(out, newImg)
	} else {
		return fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return fmt.Errorf("failed to save resized image: %v", err)
	}

	pkg.Logger.Infof("Successfully resized the image: %s", filePath)
	return nil
}

// convertFileFormat performs the file format conversion based on the provided options
func (s *FileTransformationService) convertFileFormat(filePath string, options map[string]string) error {
	// Ensure the required "format" option is provided
	newFormat, ok := options["format"]
	if !ok {
		return fmt.Errorf("missing required option: format")
	}

	// Open the existing file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Create the output file (overwrite the original)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Convert and save the image in the requested format
	switch newFormat {
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 85})
	case "png":
		err = png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported target format: %s", newFormat)
	}

	if err != nil {
		return fmt.Errorf("failed to save converted image: %v", err)
	}

	pkg.Logger.Infof("Successfully converted the image to %s format: %s", newFormat, filePath)
	return nil
}

// FileExists checks if a file exists at the given path
func (s *FileTransformationService) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}