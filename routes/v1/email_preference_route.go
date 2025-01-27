package v1

import (
    "notifications-system/internal/email_preferences/handlers"
    "notifications-system/internal/email_preferences/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupEmailPreferenceRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
    // Initialize service and handler
    emailPreferenceService := services.NewEmailPreferenceService(db)
    emailPreferenceHandler := email_preference_handler.NewEmailPreferenceHandler(emailPreferenceService)

    // Group email preference routes under /api/v1/email-preferences
    emailPreferencesGroup := v1.Group("/email-preferences")
    {
        emailPreferencesGroup.POST("/", emailPreferenceHandler.Create)
        emailPreferencesGroup.GET("/:id", emailPreferenceHandler.GetByID)
        emailPreferencesGroup.PUT("/:id", emailPreferenceHandler.Update)
        emailPreferencesGroup.DELETE("/:id", emailPreferenceHandler.Delete)
    }
}
