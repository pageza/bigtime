package main

import (
	"log"
	"net/http"

	"alchemorsel/internal/health"
)

func main() {
	http.HandleFunc("/healthz", health.Handler)
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %v", err)
	}
}
