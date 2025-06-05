package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"alchemorsel/internal/auth"
	"alchemorsel/internal/health"
	"alchemorsel/internal/profile"
	"alchemorsel/internal/recipes"
	"alchemorsel/internal/users"
)

func main() {
	store := users.NewMemoryStore()
	svc := &auth.Service{Store: store, Secret: envOr("JWT_SECRET", "secret")}
	recipeStore := recipes.NewMemoryStore()
	modStore := recipes.NewMemoryModStore()
	recipeSvc := &recipes.Service{Store: recipeStore, ModStore: modStore, LLM: recipes.FakeLLM{}}
	profileStore := profile.NewMemoryStore()
	profileSvc := &profile.Service{Store: profileStore}

	http.HandleFunc("/healthz", health.Handler)
	http.HandleFunc("/v1/users", auth.RegisterHandler(svc))
	http.HandleFunc("/v1/tokens", auth.LoginHandler(svc))
	http.HandleFunc("/v1/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			profile.GetHandler(profileSvc)(w, r)
		case http.MethodPut:
			profile.UpdateHandler(profileSvc)(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/v1/recipes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			recipes.CreateHandler(recipeSvc)(w, r)
		case http.MethodGet:
			recipes.SearchHandler(recipeSvc)(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/v1/recipes/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/modify") && r.Method == http.MethodPost {
			recipes.ModifyHandler(recipeSvc)(w, r)
			return
		}
		switch r.Method {
		case http.MethodPut:
			recipes.UpdateHandler(recipeSvc)(w, r)
		case http.MethodDelete:
			recipes.DeleteHandler(recipeSvc)(w, r)
		default:
			http.NotFound(w, r)
		}
	})

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
