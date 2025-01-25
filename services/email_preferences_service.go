// services/email_preferences_service.go
package services

import "notifications-system/models"

// Mock database (in-memory map)
var emailPreferencesDB = make(map[string]models.EmailPreferences)

func GetEmailPreferences(userID string) models.EmailPreferences {
    // Fetch from DB (mock example)
    preferences, exists := emailPreferencesDB[userID]
    if !exists {
        // Return default preferences if user doesn't exist
        return models.EmailPreferences{
            UserID:         userID,
            ProductUpdates: true,
            EventReminders: false,
        }
    }
    return preferences
}

func UpdateEmailPreferences(preferences models.EmailPreferences) {
    // Save to DB (mock example)
    emailPreferencesDB[preferences.UserID] = preferences
}