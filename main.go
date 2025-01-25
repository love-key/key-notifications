package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
	"notifications-system/database"
)

func main() {
	// Connect to PostgreSQL
    database.Connect()
    defer database.DB.Close()

    // Load .env file
    if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	} 
	// else {
	// 	log.Println(".env file loaded successfully")
	// }

    // Access environment variables
    log.Printf("Server Port: %s", /* dbHost, dbPort, */ os.Getenv("PORT"))
}