package email_preference_handler

import (
    "net/http"
    "notifications-system/internal/email_preferences/models"
    "notifications-system/internal/email_preferences/services"
    "notifications-system/internal/email_preferences/validations"
    "notifications-system/utils"  // Import utils for logging and error handling
    "github.com/gin-gonic/gin"
)

// EmailPreferenceHandler handles HTTP requests for EmailPreferences
type EmailPreferenceHandler struct {
    service *services.EmailPreferenceService
}

// NewEmailPreferenceHandler creates a new EmailPreferenceHandler instance
func NewEmailPreferenceHandler(service *services.EmailPreferenceService) *EmailPreferenceHandler {
    return &EmailPreferenceHandler{service: service}
}

// CreateEmailPreference creates a new email preference
// @Summary Create a new email preference
// @Description Creates a new email preference for the user
// @Tags email-preferences
// @Accept  json
// @Produce  json
// @Param input body models.EmailPreference true "Create Email Preference"
// @Success 201 {object} models.EmailPreference "Successfully created"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/email-preferences [post]
func (handler *EmailPreferenceHandler) Create(context *gin.Context) {
    var input models.EmailPreference

    // Call the validation function
    if err := validations.ValidateEmailPreference(&input); err != nil {
        context.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
        return
    }

    // Call the validation function
    if err := validations.ValidateEmailPreference(&input); err != nil {
        // Directly check if it's an AppError
        if appErr, ok := err.(*utils.AppError); ok {
            // Return error with message only (without code) in the response
            context.JSON(http.StatusUnprocessableEntity, gin.H{
                "statusCode": http.StatusUnprocessableEntity,
                "message":    appErr.Message, // Only include the message
            })
            return
        }
        
        // If the error is not an AppError (fallback), return a generic error
        context.JSON(http.StatusInternalServerError, gin.H{
            "statusCode": http.StatusInternalServerError,
            "message":    "Internal server error",
        })
        return
    }

    // Proceed with the service call
    if err := handler.service.Create(&input); err != nil {
        // Log error
        utils.LogError("Failed to create email preference", err)
        context.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message": "Internal server error: " + err.Error(),
		})
        return
    }

    // Log success
    utils.LogInfo("Email preference successfully created")
    // Return the created email preference
	context.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"message": "Successfully created",
	})
}

// GetByID retrieves an email preference by ID
func (handler *EmailPreferenceHandler) GetByID(context *gin.Context) {
    id := context.Param("id")
    emailPreference, err := handler.service.GetByID(id)
    if err != nil {
        // Log error
        utils.LogError("Email preference not found", err)
        context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // Log success
    utils.LogInfo("Email preference retrieved successfully")

    context.JSON(http.StatusOK, emailPreference)
}

// Update updates an email preference
func (handler *EmailPreferenceHandler) Update(context *gin.Context) {
    id := context.Param("id")
    var updatedPreference models.EmailPreference
    
    if err := context.ShouldBindJSON(&updatedPreference); err != nil {
        // Log error
        utils.LogError("Failed to bind JSON for updating email preference", err)
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := handler.service.Update(id, &updatedPreference); err != nil {
        // Log error
        utils.LogError("Failed to update email preference", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Log success
    utils.LogInfo("Email preference updated successfully")

    context.JSON(http.StatusOK, gin.H{"message": "Email preference updated successfully"})
}

// Delete deletes an email preference
func (handler *EmailPreferenceHandler) Delete(context *gin.Context) {
    id := context.Param("id")
    if err := handler.service.Delete(id); err != nil {
        // Log error
        utils.LogError("Failed to delete email preference", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Log success
    utils.LogInfo("Email preference deleted successfully")

    context.JSON(http.StatusOK, gin.H{"message": "Email preference deleted successfully"})
}
