package main

import (
	"log"
	"net/http"
	"os"

	"alchemorsel/internal/auth"
	"alchemorsel/internal/health"
	"alchemorsel/internal/recipes"
	"alchemorsel/internal/users"
)

func main() {
	store := users.NewMemoryStore()
	svc := &auth.Service{Store: store, Secret: envOr("JWT_SECRET", "secret")}
	recipeStore := recipes.NewMemoryStore()
	recipeSvc := &recipes.Service{Store: recipeStore}

	http.HandleFunc("/healthz", health.Handler)
	http.HandleFunc("/v1/users", auth.RegisterHandler(svc))
	http.HandleFunc("/v1/tokens", auth.LoginHandler(svc))
	http.HandleFunc("/v1/recipes", recipes.CreateHandler(recipeSvc))

	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %v", err)
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
