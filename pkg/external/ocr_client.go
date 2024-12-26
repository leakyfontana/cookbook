package external

import (
	"bytes"
	"cookbook/internal/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func CallOCREndpoint(payload models.OCRRequest) (string, error) {
    url := "http://localhost:8081/ocr" // URL of the OCR endpoint

    // Create the request payload
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    // Make the HTTP POST request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Check for non-200 status code
    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to call OCR endpoint")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}