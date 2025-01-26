package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // Import gin-swagger
	"github.com/swaggo/files"       // Import swagger files

	"notifications-system/config"
	"notifications-system/database"
	_ "notifications-system/docs" // Import the generated docs
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Connect to the database
	database.Connect() // No need to store the return value since we're using the global DB
	defer database.Close() // Ensure we close the connection when the application terminates

	// Handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Set Gin to Release Mode (disable debug logging)
	gin.SetMode(gin.ReleaseMode)
	
	// Create Gin router
	r := gin.New()

	// Serve Swagger UI at the /swagger endpoint (uses pre-generated Swagger docs)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Your other routes go here...
	// Example: r.GET("/api/v1/users", getUsersHandler)

	// Start the server on the specified port
	log.Printf("Server Port: %s", config.GetEnv("PORT", "3000"))
	go func() {
		if err := r.Run(":" + config.GetEnv("PORT", "3000")); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for shutdown signal (only one time)
	<-ctx.Done()

	// Shutdown logic
	log.Println("Received termination signal, shutting down...")
	time.Sleep(1 * time.Second) // Simulate cleanup if needed
	log.Println("Shutdown complete")
}
