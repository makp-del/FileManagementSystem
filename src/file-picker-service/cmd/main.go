package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"file-picker-service/config"
	"file-picker-service/internal/db"
	"file-picker-service/internal/handlers"
	"file-picker-service/internal/pkg"
	"file-picker-service/internal/services"
)

func main() {
	// Load configuration (environment variables, etc.)
	cfg := config.LoadConfig()

	// Initialize the logger
	pkg.InitLogger()

	// Initialize the database connection
	db.InitDatabase()


	// Initialize gRPC clients for communication with other services

	permissionsClient, err := services.NewPermissionClient(cfg.GrpcAddresses.PermissionsAddress)

	if err != nil {
		log.Fatalf("Failed to create permissions client: %v", err)
	}

	downloaderClient, err := services.NewDownloaderClient(cfg.GrpcAddresses.FileDownloaderAddress)
	if err != nil {
		log.Fatalf("Failed to create downloader client: %v", err)
	}

	transformationConn, err := grpc.Dial(cfg.GrpcAddresses.TransformationsAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Transformation Service: %v", err)
	}
	defer transformationConn.Close()
	transformationClient := services.NewTransformationClient(transformationConn)

	notificationConn, err := grpc.Dial(cfg.GrpcAddresses.NotificationAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Notification Service: %v", err)
	}
	defer notificationConn.Close()
	notificationClient := services.NewNotificationClient(notificationConn)

	// Set up the Gin router
	router := gin.Default()

	// Register routes and handlers
	router.POST("/api/upload", handlers.FileUploadHandler(permissionsClient))
	router.GET("/api/files", handlers.ListFilesHandler(permissionsClient))
	router.POST("/api/files/download", handlers.FileDownloadHandler(permissionsClient, downloaderClient))
	router.POST("/api/transform/:id", handlers.FileTransformationHandler(permissionsClient, transformationClient, notificationClient))
	router.POST("/api/files/share", handlers.AddPermissionsHandler(permissionsClient))

	// Start the server
	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}
	log.Printf("File-Picker-Service is running on port %s", port)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
