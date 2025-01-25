package models

type EmailPreferences struct {
    UserID         string `json:"user_id"`
    ProductUpdates bool   `json:"product_updates"`
    EventReminders bool   `json:"event_reminders"`
}