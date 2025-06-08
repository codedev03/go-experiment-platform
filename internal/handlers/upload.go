package handlers

import (
    "net/http"
    "goexperiment/internal/storage"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB limit

    file, fileHeader, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to read file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    url, err := storage.UploadToS3(file, fileHeader)
    if err != nil {
        http.Error(w, "Upload to S3 failed: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(url))
}
