package v1

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupV1Routes(router *gin.Engine, db *gorm.DB) {
    // Group all routes under /api/v1
    v1 := router.Group("/api/v1")
    {
        // Setup email preference routes
        SetupEmailPreferenceRoutes(v1, db)
        
        // You can define other routes here similarly
        // SetupNotificationPreferenceRoutes(v1, db)
    }
}
