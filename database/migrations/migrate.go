package main

import (
    "fmt"
    "log"
	"notifications-system/database"      // for database connection    
	"notifications-system/internal/email_preferences/models"
    "gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
    // Run migrations for all models
    fmt.Println("Running migrations...")
    err := db.AutoMigrate(
		&models.EmailPreference{},
		// &models.User{},  // to add other models for migrations
	)
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }
    fmt.Println("Migration completed successfully!")
}

func main() {
    // Establish database connection
    db := database.Connect()
    defer database.Close()

    // Run the migrations
    runMigrations(db)
}

