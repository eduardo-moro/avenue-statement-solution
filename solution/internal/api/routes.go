package api

import "net/http"

// RegisterRoutes registers the API routes.
func RegisterRoutes() {
	http.HandleFunc("/events", EventsHandler)
	http.HandleFunc("/statement/", StatementHandler)
}