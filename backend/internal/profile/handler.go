package profile

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetHandler handles GET /v1/profile.
func GetHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := s.Get(r.Context(), 1)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.NotFound(w, r)
			} else {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}
		_ = json.NewEncoder(w).Encode(p)
	}
}

// UpdateHandler handles PUT /v1/profile.
func UpdateHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		p, err := s.Update(r.Context(), 1, req)
		if err != nil {
			if errors.Is(err, ErrInvalidInput) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}
		_ = json.NewEncoder(w).Encode(p)
	}
}
