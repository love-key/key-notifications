package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"notifications-system/config"
	"notifications-system/database"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Connect to the database
	database.Connect() // No need to store the return value since we're using the global DB
	defer database.Close()  // Ensure we close the connection when the application terminates

	// Handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start application logic
	log.Printf("Server Port: %s", config.GetEnv("PORT", "3000"))

	// Wait for shutdown signal (only one time)
	<-ctx.Done()

	// Shutdown logic
	log.Println("Received termination signal, shutting down...")
	time.Sleep(1 * time.Second) // Simulate cleanup if needed
	log.Println("Shutdown complete")
}
