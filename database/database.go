package database

import (
    "context"
    "fmt"
    "os"
	"log"
    "time"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Build the connection string
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("PG_HOST"),
        os.Getenv("PG_PORT"),
        os.Getenv("PG_USER"),
        os.Getenv("PG_PASSWORD"),
        os.Getenv("PG_DB"),
    )

    // Create a connection pool
    config, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        log.Fatalf("Unable to parse database config: %v", err)
    }

    // Configure connection pool
    config.MaxConns = 10
    config.MinConns = 2
    config.MaxConnLifetime = time.Hour
    config.MaxConnIdleTime = time.Minute * 30

    // Establish the connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    DB, err = pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    log.Println("Connected to database:", os.Getenv("PG_DB"),)
}