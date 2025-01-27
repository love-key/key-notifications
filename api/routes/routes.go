// api/routes/routes.go
package routes

import (
    "github.com/gorilla/mux"
    "notifications-system/api/handlers"
)

func SetupRoutes(r *mux.Router) {
    // Email Preferences routes
    r.HandleFunc("/email-preferences", handlers.GetEmailPreferencesHandler).Methods("GET")
    r.HandleFunc("/email-preferences", handlers.UpdateEmailPreferencesHandler).Methods("PUT")
}