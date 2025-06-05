package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	store := NewMemoryStore()
	svc := &Service{Store: store, Secret: []byte("secret")}

	reg := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"email": "a@x.com", "password": "p"})
	RegisterHandler(svc)(reg, httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body)))
	if reg.Code != http.StatusCreated {
		t.Fatalf("reg code %d", reg.Code)
	}

	login := httptest.NewRecorder()
	LoginHandler(svc)(login, httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body)))
	if login.Code != http.StatusOK {
		t.Fatalf("login code %d", login.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(login.Body).Decode(&resp); err != nil || resp["token"] == "" {
		t.Fatalf("bad token resp: %v", err)
	}
}
