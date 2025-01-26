package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // Import gin-swagger
	"github.com/swaggo/files"       // Import swagger files

	"notifications-system/config"
	"notifications-system/database"
	"notifications-system/routes/v1" // Import v1 routes
	_ "notifications-system/docs" // Import the generated docs
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Connect to the database
	db := database.Connect() // No need to store the return value since we're using the global DB
	defer database.Close() // Ensure we close the connection when the application terminates

	// Handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Set Gin to Release Mode (disable debug logging)
	gin.SetMode(gin.ReleaseMode)

	// Create a single Gin router
	router := gin.New()

	// Simple route for testing
	router.GET("/", func(c *gin.Context) {
		// log.Println("Hello, World! route hit")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Serve Swagger UI at the /swagger endpoint (uses pre-generated Swagger docs)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup versioned routes (v1)
	v1.SetupV1Routes(router, db)  // Pass Gin router and DB connection to setup the routes)

	// Start the server on the specified port
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server Port: %s", port)
	go func() {
		if err := router.Run(":" + port); err != nil {
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