package utils

import "fmt"

// Custom error codes
const (
    ErrCodeNotFound     = "NOT_FOUND"
    ErrCodeBadRequest   = "BAD_REQUEST"
    ErrCodeServerError  = "SERVER_ERROR"
)

// AppError represents a custom error type with code, message, and additional context
type AppError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

// NewAppError creates a new custom error instance
func NewAppError(code, message string) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
    }
}

// Error implements the error interface
func (e *AppError) Error() string {
    return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}
