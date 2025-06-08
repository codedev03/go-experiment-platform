package handlers

import (
    "encoding/json"
    "net/http"

    "goexperiment/internal/db"
    "goexperiment/internal/model"
)

func CreateExperiment(w http.ResponseWriter, r *http.Request) {
    var exp model.Experiment
    if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    query := "INSERT INTO experiments (name, description, status) VALUES (?, ?, ?)"
    _, err := db.DB.Exec(query, exp.Name, exp.Description, exp.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Experiment created"})
}

func ListExperiments(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, name, description, status FROM experiments")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var experiments []model.Experiment
    for rows.Next() {
        var exp model.Experiment
        if err := rows.Scan(&exp.ID, &exp.Name, &exp.Description, &exp.Status); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        experiments = append(experiments, exp)
    }

    json.NewEncoder(w).Encode(experiments)
}
