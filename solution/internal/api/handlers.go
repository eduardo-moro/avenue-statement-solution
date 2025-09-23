package api

import "net/http"

// EventsHandler handles the /events endpoint.
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// StatementHandler handles the /statement endpoint.
func StatementHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
