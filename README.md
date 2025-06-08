# Go Experiment Platform 🎯

A lightweight backend service for managing A/B experiments, built with Go and MySQL.

## 📌 Features

- ✅ Create and list A/B experiments via RESTful APIs  
- ✅ Clean folder structure using Go best practices (`internal/`, `model/`, etc.)  
- ✅ MySQL database integration using Go's `database/sql` package  
- ✅ Input validation and structured JSON responses  
- ✅ Easy to test with Postman or curl

## 🗂 Project Structure
goexperiment/
├── go.mod # Go module file
├── main.go # Application entry point
├── internal/
│ ├── db/ # Database connection logic
│ │ └── db.go
│ └── handlers/ # HTTP route handlers
│ └── handler.go
├ ── model/
│  └── experiment.go # Experiment struct/model
