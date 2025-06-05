package main

import (
	"log"
	"net/http"

	"alchemorsel/internal/health"
	"alchemorsel/internal/user"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health.Handler)

	store := user.NewMemoryStore()
	svc := &user.Service{Store: store}
	mux.Handle("/v1/users", &user.Handler{Service: svc})

	server := &http.Server{Addr: ":8080", Handler: mux}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %v", err)
	}
}
