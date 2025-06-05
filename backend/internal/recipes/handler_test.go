package recipes

import (
	"bytes"

	"context"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSearchHandler(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	if _, err := svc.Create(context.Background(), 1, CreateRequest{Title: "Carrot Soup", Ingredients: []string{"carrot"}, Steps: []string{"mix"}}); err != nil {
		t.Fatalf("create: %v", err)
	}
	r := httptest.NewRequest(http.MethodGet, "/v1/recipes?q=carrot", nil)
	w := httptest.NewRecorder()
	SearchHandler(svc)(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("code %d", w.Code)
	}
	var resp []Recipe
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil || len(resp) != 1 {
		t.Fatalf("bad response: %v", err)
	}
}

func TestModifyHandler(t *testing.T) {
	svc := &Service{Store: NewMemoryStore(), ModStore: NewMemoryModStore(), LLM: FakeLLM{}}
	rOrig, _ := svc.Create(context.Background(), 1, CreateRequest{Title: "Stew", Ingredients: []string{"beef"}, Steps: []string{"cook"}})
	body, _ := json.Marshal(map[string]string{"prompt": "spicy"})
	req := httptest.NewRequest(http.MethodPost, "/v1/recipes/"+strconv.FormatInt(rOrig.ID, 10)+"/modify", bytes.NewReader(body))
	w := httptest.NewRecorder()
	ModifyHandler(svc)(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("code %d", w.Code)
	}
}

func TestModifyHandler_Bad(t *testing.T) {
	svc := &Service{Store: NewMemoryStore(), ModStore: NewMemoryModStore(), LLM: FakeLLM{}}
	body, _ := json.Marshal(map[string]string{"prompt": "fail"})
	req := httptest.NewRequest(http.MethodPost, "/v1/recipes/1/modify", bytes.NewReader(body))
	w := httptest.NewRecorder()
	ModifyHandler(svc)(w, req)
	if w.Code == http.StatusCreated {
		t.Fatalf("expected error")
	}
}

func TestUpdateHandler(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	rOrig, _ := svc.Create(context.Background(), 1, CreateRequest{Title: "Stew", Ingredients: []string{"beef"}, Steps: []string{"cook"}})
	body, _ := json.Marshal(CreateRequest{Title: "Veg Stew", Ingredients: []string{"beef"}, Steps: []string{"cook"}})
	req := httptest.NewRequest(http.MethodPut, "/v1/recipes/"+strconv.FormatInt(rOrig.ID, 10), bytes.NewReader(body))
	w := httptest.NewRecorder()
	UpdateHandler(svc)(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("code %d", w.Code)
	}
}

func TestDeleteHandler(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	rOrig, _ := svc.Create(context.Background(), 1, CreateRequest{Title: "Stew", Ingredients: []string{"beef"}, Steps: []string{"cook"}})
	req := httptest.NewRequest(http.MethodDelete, "/v1/recipes/"+strconv.FormatInt(rOrig.ID, 10), nil)
	w := httptest.NewRecorder()
	DeleteHandler(svc)(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("code %d", w.Code)
	}
}
