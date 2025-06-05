package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ServeHTTP(t *testing.T) {
	store := NewMemoryStore()
	svc := &Service{Store: store}
	h := &Handler{Service: svc}

	body, _ := json.Marshal(RegistrationRequest{
		Email:    "foo@example.com",
		Password: "secret",
	})
	req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req.WithContext(context.Background()))

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rr.Code)
	}
	var res RegistrationResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Email != "foo@example.com" {
		t.Fatalf("unexpected email %s", res.Email)
	}
}
