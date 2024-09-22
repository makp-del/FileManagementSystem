package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"auth-service/config"
	"auth-service/internal/db"
	"auth-service/internal/handlers"
	"auth-service/internal/models"
	"auth-service/pkg"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Initialize the logger
	pkg.InitLogger()

	// Load environment configuration
	cfg := config.LoadConfig()

	// Initialize the database
	db.InitDB()

	// Perform auto-migration for the User model
	db.DB.AutoMigrate(&models.User{})

	// Create default admin user if not exists
	createDefaultAdminUser()

	// Load private key for JWT
	err := pkg.LoadPrivateKey(cfg.PrivateKeyPath)
	if err != nil {
		pkg.Logger.Fatal("Error loading private key: ", err)
	}

	// Setup Gin router
	r := gin.Default()

	// Routes
	r.POST("/api/login", handlers.Login)
	r.POST("/api/register", handlers.Register)

	// Start the server with graceful shutdown
	srv := startServer(r)

	// Handle graceful shutdown on interrupt signal
	gracefulShutdown(srv)
}

// createDefaultAdminUser creates a default admin user if it doesn't already exist
func createDefaultAdminUser() {
	var adminUser models.User
	if err := models.GetUserByUsername(db.DB, "admin", &adminUser); err != nil {
		// Hash the password for the admin user
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)

		// Create the admin user
		adminUser := models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Email:    "admin@example.com",
		}

		if err := models.CreateUser(db.DB, &adminUser); err != nil {
			fmt.Println("Error creating admin user:", err)
		} else {
			fmt.Println("Admin user created successfully.")
		}
	}
	
}

// startServer starts the Gin HTTP server in a separate goroutine
func startServer(r *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start the server in a goroutine so that it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			pkg.Logger.Fatalf("Error starting server: %s", err)
		}
	}()

	pkg.Logger.Info("Server started and running on port 8080")
	return srv
}

// gracefulShutdown gracefully shuts down the server when an interrupt signal is received
func gracefulShutdown(srv *http.Server) {
	// Create a channel to listen for system interrupts
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Block until an interrupt signal is received
	<-quit
	pkg.Logger.Info("Received shutdown signal, shutting down server...")

	// Create a deadline for the server shutdown (e.g., 5 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully, waiting for ongoing requests to complete
	if err := srv.Shutdown(ctx); err != nil {
		pkg.Logger.Fatalf("Server forced to shutdown: %s", err)
	}

	pkg.Logger.Info("Server shut down gracefully")
}