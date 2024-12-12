package main

import (
	"cookbook/config"
	"cookbook/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize routes
	http.HandleFunc("/ocr", handlers.OCRHandler)
	http.HandleFunc("/recipes", handlers.RecipeHandler)
	http.HandleFunc("/google-docs", handlers.GoogleDocsHandler)

	// Start server
	log.Printf("Server running at %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}
