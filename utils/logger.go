package utils

import (
    "log"
    "os"
)

// Logger setup
var logger *log.Logger

func init() {
    logger = log.New(os.Stdout, "[APP_LOG] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages
func LogInfo(message string) {
    logger.Println("[INFO]", message)
}

// LogError logs error messages
func LogError(message string, err error) {
    logger.Println("[ERROR]", message, err)
}

// LogWarning logs warning messages
func LogWarning(message string) {
    logger.Println("[WARNING]", message)
}
