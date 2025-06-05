package main

import (
	"log"
	"net/http"

	"alchemorsel/internal/health"
	"alchemorsel/internal/user"
)

func main() {
	http.HandleFunc("/healthz", health.Handler)
	store := user.NewMemoryStore()
	svc := &user.Service{Store: store, Secret: []byte("secret")}
	http.HandleFunc("/v1/users", user.RegisterHandler(svc))
	http.HandleFunc("/v1/tokens", user.LoginHandler(svc))
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %v", err)
	}
}
