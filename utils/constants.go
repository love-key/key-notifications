package utils

// Constants for API versioning
const (
    APIVersionV1 = "/api/v1"
)

// HTTP Status Codes for easy reference
const (
    StatusOK                  = 200
    StatusCreated             = 201
    StatusBadRequest          = 400
    StatusUnauthorized        = 401
    StatusForbidden           = 403
    StatusNotFound            = 404
	StatusUnprocessableEntity = 422
    StatusInternalServerError = 500

)

// Environment variables
const (
    EnvProduction  = "production"
    EnvDevelopment = "development"
    EnvTesting     = "testing"
)
