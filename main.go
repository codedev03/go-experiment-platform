package main

import (
    "log"
    "net/http"
    "goexperiment/internal/db"
    "goexperiment/internal/handlers"
)

func main() {
    db.InitDB()

    http.HandleFunc("/experiments", handlers.ListExperiments)
    http.HandleFunc("/create", handlers.CreateExperiment)

    log.Println("✅ Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
