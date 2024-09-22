package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"file-downloader-service/config"
	"file-downloader-service/internal/handlers"
	"file-downloader-service/internal/pkg"
	"file-downloader-service/proto/generated/filedownloader"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize logger
	pkg.InitLogger()

	// Load configuration
	err := config.LoadConfig()
	if err != nil {
		pkg.Logger.Fatalf("Failed to load configuration: %v", err)
	}

	// Start the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.AppConfig.Port))
	if err != nil {
		pkg.Logger.Fatalf("Failed to listen on port %s: %v", config.AppConfig.Port, err)
	}

	grpcServer := grpc.NewServer()

	// Register the FileDownloaderService
	filedownloader.RegisterFileDownloaderServiceServer(grpcServer, &handlers.FileDownloaderHandler{})

	// Enable gRPC reflection for easier development and debugging
	reflection.Register(grpcServer)

	// Run the server in a goroutine so we can handle graceful shutdown
	go func() {
		pkg.Logger.Infof("File Downloader Service is running on port %s", config.AppConfig.Port)
		if err := grpcServer.Serve(lis); err != nil {
			pkg.Logger.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	pkg.Logger.Info("Shutting down File Downloader Service...")

	grpcServer.GracefulStop()
	pkg.Logger.Info("File Downloader Service stopped gracefully")
}