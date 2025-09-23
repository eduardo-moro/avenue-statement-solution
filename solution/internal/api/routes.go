package api

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/events", EventsHandler)
	http.HandleFunc("/statement/", StatementHandler)
}
