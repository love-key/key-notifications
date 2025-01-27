package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// DBConfig holds the database configuration values
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Dialect  string
}

// LoadEnv loads the .env file into environment variables
func LoadEnv() error {
	// Try loading the .env file
	err := godotenv.Load()
	if err != nil {
		return err
	}
	// log.Println(".env file loaded successfully")
	return nil
}

// LoadDBConfig loads the database configuration from environment variables
func LoadDBConfig() *DBConfig {
	// Fetch the database-related environment variables
	dbConfig := &DBConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   os.Getenv("PG_DB"),
		Dialect:  os.Getenv("DIALECT"),
	}

	// Check if any of the required database configuration values are missing
	if dbConfig.Host == "" || dbConfig.Port == "" || dbConfig.User == "" || dbConfig.Password == "" || dbConfig.DBName == "" {
		log.Fatal("Database configuration is missing required values")
	}

	// log.Println("Database configuration loaded successfully")
	return dbConfig
}

// GetEnv fetches an environment variable or returns a default value if not set
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
