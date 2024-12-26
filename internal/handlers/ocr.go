package handlers

import (
	"cookbook/internal/models"
	"cookbook/pkg/external"
	"encoding/json"
	"net/http"
)

func OCRHandler(w http.ResponseWriter, r *http.Request) {

    var request models.OCRRequest

    // Extract image path from request (this is just an example, adjust as needed)
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }
    imagePath := string(request.ImagePath)
    if imagePath == "" {
        http.Error(w, "Image path is required", http.StatusBadRequest)
        return
    }

    // Call the OCR endpoint
    text, err := external.CallOCREndpoint(request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the OCR result
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(text))
}
