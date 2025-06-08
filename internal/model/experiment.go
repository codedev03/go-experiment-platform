package model

type Experiment struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Status      string `json:"status"` // "running", "paused", "completed"
}
