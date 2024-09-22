package main

import (
	"log"
	"net"
	"net/http"
	"notification-service/config"
	"notification-service/internal/handlers"
	"notification-service/utils/logger"
	"notification-service/internal/services"
	"notification-service/internal/websocket"
	"notification-service/proto/generated/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize the logger
	logger.Init()

	// Load configuration
	cfg := config.LoadConfig()

	// Create a new WebSocket hub
	hub := websocket.NewHub()

	// Create a new NotificationService
	notificationService := services.NewNotificationService(hub)

	// Set up the WebSocket handler
	webSocketHandler := &handlers.WebSocketHandler{
		Hub: hub,
	}

	// Run the WebSocket hub
	go hub.Run()

	// Start the HTTP server for WebSocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		webSocketHandler.ServeWebSocket(w, r)
	})

	// Set up the gRPC server
	grpcServer := grpc.NewServer()

	// Register the NotificationHandler (gRPC service) with the server
	notification.RegisterNotificationServiceServer(grpcServer, &handlers.NotificationHandler{
		NotificationService: notificationService,
	})

	// Enable reflection for debugging with grpcurl
	reflection.Register(grpcServer)

	// Start the gRPC server in a separate goroutine
	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port %s: %v", cfg.GRPCPort, err)
		}

		log.Printf("Starting gRPC server on port %s...\n", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start the WebSocket server
	log.Printf("Starting WebSocket server on port %s...\n", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start WebSocket server: %v", err)
	}
}