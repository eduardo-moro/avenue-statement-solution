package main

import (
	"avenue-golang/solution/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	api.RegisterRoutes()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	log.Println("🧠 Iniciando servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("❌ Falha ao iniciar servidor: %v", err)
	}
}
