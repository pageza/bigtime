package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"alchemorsel/internal/users"
)

func setup() *Service {
	return &Service{Store: users.NewMemoryStore(), Secret: "test"}
}

func TestRegisterHandler(t *testing.T) {
	svc := setup()
	cases := []struct {
		name       string
		req        RegisterRequest
		wantStatus int
	}{
		{"ok", RegisterRequest{Email: "a@b.com", Password: "password"}, http.StatusCreated},
		{"weak", RegisterRequest{Email: "x@y.com", Password: "short"}, http.StatusBadRequest},
		{"dup", RegisterRequest{Email: "a@b.com", Password: "password"}, http.StatusConflict},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.req)
			r := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewReader(body))
			w := httptest.NewRecorder()
			RegisterHandler(svc)(w, r)
			if w.Code != tc.wantStatus {
				t.Fatalf("expected %d, got %d", tc.wantStatus, w.Code)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	svc := setup()
	// register a user first
	_, err := svc.Register(context.Background(), RegisterRequest{Email: "login@t.com", Password: "password"})
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name       string
		req        LoginRequest
		wantStatus int
	}{
		{"ok", LoginRequest{Email: "login@t.com", Password: "password"}, http.StatusOK},
		{"badpass", LoginRequest{Email: "login@t.com", Password: "wrong"}, http.StatusUnauthorized},
		{"nouser", LoginRequest{Email: "missing@t.com", Password: "pass"}, http.StatusUnauthorized},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.req)
			r := httptest.NewRequest(http.MethodPost, "/v1/tokens", bytes.NewReader(body))
			w := httptest.NewRecorder()
			LoginHandler(svc)(w, r)
			if w.Code != tc.wantStatus {
				t.Fatalf("expected %d, got %d", tc.wantStatus, w.Code)
			}
		})
	}
}
