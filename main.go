package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "notifications-system/api/routes"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Set up routes
    routes.SetupRoutes(r)

    // Start the server
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Server error:", err)
    }
}