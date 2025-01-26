package main

import (
    "fmt"
    "log"
    "notifications-system/database"
    "notifications-system/database/models"

    "gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
    // Run migrations for all models
    fmt.Println("Running migrations...")
    err := db.AutoMigrate(&models.EmailPreference{})
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

