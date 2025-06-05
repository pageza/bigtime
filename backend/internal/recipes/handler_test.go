package recipes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	body, _ := json.Marshal(CreateRequest{Title: "T", Ingredients: []string{"i"}, Steps: []string{"s"}})
	r := httptest.NewRequest(http.MethodPost, "/v1/recipes", bytes.NewReader(body))
	w := httptest.NewRecorder()
	CreateHandler(svc)(w, r)
	if w.Code != http.StatusCreated {
		t.Fatalf("code %d", w.Code)
	}
	var resp Recipe
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil || resp.ID == 0 {
		t.Fatalf("bad response: %v", err)
	}
}

func TestCreateHandler_BadInput(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	body, _ := json.Marshal(CreateRequest{})
	r := httptest.NewRequest(http.MethodPost, "/v1/recipes", bytes.NewReader(body))
	w := httptest.NewRecorder()
	CreateHandler(svc)(w, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("code %d", w.Code)
	}
}
