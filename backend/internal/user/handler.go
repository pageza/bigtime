package user

import (
	"encoding/json"
	"errors"
	"net/http"
)

// RegistrationRequest represents a user sign-up request.
type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegistrationResponse represents a successful registration.
type RegistrationResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

// Handler processes user registration requests.
type Handler struct {
	Service *Service
}

// ServeHTTP registers a new user and returns a JSON response.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	user, err := h.Service.Register(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, ErrEmailExists) {
			http.Error(w, "email already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	res := RegistrationResponse{ID: user.ID, Email: user.Email}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
