package main

import (
	"fmt"
	"log"
	"net"

	"permission-service/config"
	"permission-service/internal/db"
	"permission-service/internal/handlers"
	"permission-service/utils/logger"
	"permission-service/internal/services"
	"permission-service/proto/generated/permission"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize the logger
	logger.Init()

	// Load the configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	err := db.InitDB(cfg.DBUrl) // Pass DB URL from the config
	if err != nil {
		logger.Error.Println("Failed to initialize database:", err)
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Initialize the permission service
	permissionService := services.NewPermissionService(db.DBConn)

	// Set up the gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		logger.Error.Println("Failed to listen on port", cfg.Port, "Error:", err)
		log.Fatalf("Failed to listen on port %s: %v", cfg.Port, err)
	}

	grpcServer := grpc.NewServer()

	// Register the PermissionService with the gRPC server
	permission.RegisterPermissionServiceServer(grpcServer, &handlers.PermissionHandler{
		PermissionService: permissionService,
	})

	// Enable reflection for tools like grpcurl
	reflection.Register(grpcServer)

	logger.Info.Println("Permission Service is running on port", cfg.Port)

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		logger.Error.Println("Failed to serve gRPC server:", err)
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}