package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "backend/internal/config"
    "backend/internal/database"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize the database
    db, err := database.InitDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    // Set up the router
    router := mux.NewRouter()

    // Define your routes here
    // e.g., router.HandleFunc("/api/v1/example", ExampleHandler).Methods("GET")

    // Start the server
    log.Printf("Server is running on %s", cfg.Port)
    log.Fatal(http.ListenAndServe(cfg.Port, router))
}