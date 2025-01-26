package main

import (
    "fmt"
    "log"
    "notifications-system/database"
    "gorm.io/gorm"
)

func rollbackMigrations(db *gorm.DB) {
    // Example of rolling back migrations (dropping a table)
    fmt.Println("Rolling back migrations...")
    err := db.Migrator().DropTable("emailPreferences")
    if err != nil {
        log.Fatalf("Rollback failed: %v", err)
    }
    fmt.Println("Rollback completed successfully!")
}

func main() {
    // Establish database connection
    db := database.Connect()
    defer database.Close()

    // Rollback migrations
    rollbackMigrations(db)
}