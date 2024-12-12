package handlers

import "net/http"

func OCRHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for OCR logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OCR Endpoint"))
}
