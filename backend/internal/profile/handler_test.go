package profile

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}

	// Update profile
	body, _ := json.Marshal(UpdateRequest{DisplayName: "Bob"})
	req := httptest.NewRequest(http.MethodPut, "/v1/profile", bytes.NewReader(body))
	w := httptest.NewRecorder()
	UpdateHandler(svc)(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("update code %d", w.Code)
	}

	// Get profile
	g := httptest.NewRecorder()
	GetHandler(svc)(g, httptest.NewRequest(http.MethodGet, "/v1/profile", nil))
	if g.Code != http.StatusOK {
		t.Fatalf("get code %d", g.Code)
	}
	var p Profile
	if err := json.NewDecoder(g.Body).Decode(&p); err != nil || p.DisplayName != "Bob" {
		t.Fatalf("bad body: %v", err)
	}
}
