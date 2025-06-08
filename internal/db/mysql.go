package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // global DB variable

func InitDB() {
    dsn := "root:yourpassword@tcp(127.0.0.1:3306)/goexperiment_db"
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error opening DB:", err)
        return
    }

    if err := DB.Ping(); err != nil {
        fmt.Println("Ping error:", err)
        return
    }

    fmt.Println("✅ Successfully connected to MySQL!")
}
