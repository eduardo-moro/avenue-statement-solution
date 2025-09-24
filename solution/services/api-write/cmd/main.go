package main

import (
	"log"
	"net/http"
	"os"

	"avenue-golang/services/api-write/internal/handlers"
	"avenue-golang/services/api-write/internal/services"
)

func main() {
	svc := services.NewService()
	h := handlers.NewHandler(svc)

	http.HandleFunc("/health", h.HealthCheck)
	http.HandleFunc("/events", h.EventsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API Write Service listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
