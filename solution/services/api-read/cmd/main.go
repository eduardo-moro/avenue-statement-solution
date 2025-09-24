package main

import (
	"log"
	"net/http"
	"os"

	"avenue-golang/services/api-read/internal/handlers"
	"avenue-golang/services/api-read/internal/services"
)

func main() {
	svc := services.NewService()
	h := handlers.NewHandler(svc)

	http.HandleFunc("/health", h.HealthCheck)
	http.HandleFunc("/statement/", h.StatementHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API Read Service listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
