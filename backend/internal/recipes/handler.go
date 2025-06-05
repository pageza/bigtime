package recipes

import (
	"encoding/json"
	"errors"
	"net/http"

	"strconv"

)

// CreateHandler returns an HTTP handler for POST /v1/recipes.
func CreateHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		recipe, err := s.Create(r.Context(), 1, req)
		if err != nil {
			if errors.Is(err, ErrInvalidInput) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(recipe)
	}
}


// SearchHandler returns an HTTP handler for GET /v1/recipes.
func SearchHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		limit, _ := strconv.Atoi(q.Get("limit"))
		req := SearchRequest{
			Q:          q.Get("q"),
			Ingredient: q["ingredient"],
			Tag:        q["tag"],
			Page:       page,
			Limit:      limit,
		}
		recipes, err := s.Search(r.Context(), req)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(recipes)
	}
}

