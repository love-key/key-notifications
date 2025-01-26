package database

import (
	"log"
	"time"
	"notifications-system/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB holds the GORM database connection
var DB *gorm.DB

// Connect establishes the database connection and returns the underlying *sql.DB
func Connect() *gorm.DB {
	// Load environment variables and configurations
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	dbConfig := config.LoadDBConfig()

	// Build DSN (Data Source Name)
	dsn := "host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DBName +
		" sslmode=disable"

	// Initialize GORM with the appropriate dialect
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve SQL DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10) // Automatically close idle connections after 10 minutes

	log.Println("Database connection established successfully")
	return DB
}

// Close closes the database connection
func Close() {
	// Retrieve the SQL DB and close the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve SQL DB for closing: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
	log.Println("Database connection closed successfully")
}
