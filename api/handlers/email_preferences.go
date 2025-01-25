// api/handlers/email_preferences.go
package handlers

import (
    "encoding/json"
    "net/http"
    "notifications-system/models"
    "notifications-system/services"
)

func GetEmailPreferencesHandler(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from query parameters
    userID := r.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(w, "user_id is required", http.StatusBadRequest)
        return
    }

    // Fetch preferences from the service
    preferences := services.GetEmailPreferences(userID)

    // Return preferences as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(preferences)
}

func UpdateEmailPreferencesHandler(w http.ResponseWriter, r *http.Request) {
    // Decode the request body into EmailPreferences struct
    var preferences models.EmailPreferences
    if err := json.NewDecoder(r.Body).Decode(&preferences); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Validate user ID
    if preferences.UserID == "" {
        http.Error(w, "user_id is required", http.StatusBadRequest)
        return
    }

    // Update preferences using the service
    services.UpdateEmailPreferences(preferences)

    // Return success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Email preferences updated!"))
}