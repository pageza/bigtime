package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"alchemorsel/internal/users"
)

// Service handles user registration and login.
type Service struct {
	Store  users.Store
	Secret string
}

// RegisterRequest holds input for registration.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest holds input for login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register registers a new user and returns a JWT.
func (s *Service) Register(ctx context.Context, req RegisterRequest) (string, error) {
	if len(req.Password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}
	hashed, err := hashPassword([]byte(req.Password))
	if err != nil {
		return "", err
	}
	u := &users.User{Email: req.Email, PasswordHash: hashed}
	if err := s.Store.Create(ctx, u); err != nil {
		return "", err
	}
	return generateJWT(u.ID, s.Secret), nil
}

// Login authenticates a user and returns a JWT.
func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	u, err := s.Store.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	if !verifyPassword([]byte(req.Password), u.PasswordHash) {
		return "", errors.New("invalid credentials")
	}
	return generateJWT(u.ID, s.Secret), nil
}

func generateJWT(userID int64, secret string) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payloadBytes, _ := json.Marshal(map[string]any{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	signing := header + "." + payload
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signing))
	sig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return signing + "." + sig
}

// RegisterHandler handles POST /v1/users
func RegisterHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		token, err := s.Register(r.Context(), req)
		if err != nil {
			status := http.StatusBadRequest
			if errors.Is(err, users.ErrEmailExists) {
				status = http.StatusConflict
			}
			http.Error(w, err.Error(), status)
			return
		}
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(token)); err != nil {
			http.Error(w, "write error", http.StatusInternalServerError)
		}
	}
}

// LoginHandler handles POST /v1/tokens
func LoginHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		token, err := s.Login(r.Context(), req)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if _, err := w.Write([]byte(token)); err != nil {
			http.Error(w, "write error", http.StatusInternalServerError)
		}
	}
}
