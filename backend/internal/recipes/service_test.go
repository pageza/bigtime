package recipes

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestService_Create(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	ctx := context.Background()
	cases := []struct {
		name    string
		req     CreateRequest
		wantErr bool
	}{
		{"ok", CreateRequest{Title: "T", Ingredients: []string{"i"}, Steps: []string{"s"}}, false},
		{"missing", CreateRequest{Ingredients: []string{"i"}, Steps: []string{"s"}}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := svc.Create(ctx, 1, tc.req)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected: %v", err)
			}
		})
	}
}

func TestService_Search(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	ctx := context.Background()
	if _, err := svc.Create(ctx, 1, CreateRequest{Title: "Apple Pie", Ingredients: []string{"apple", "flour"}, Steps: []string{"mix"}}); err != nil {
		t.Fatalf("create apple: %v", err)
	}
	if _, err := svc.Create(ctx, 1, CreateRequest{Title: "Banana Bread", Ingredients: []string{"banana"}, Steps: []string{"mix"}, Tags: []string{"dessert"}}); err != nil {
		t.Fatalf("create banana: %v", err)
	}
	results, err := svc.Search(ctx, SearchRequest{Q: "banana", Page: 1, Limit: 10})
	if err != nil || len(results) != 1 || !strings.Contains(results[0].Title, "Banana") {
		t.Fatalf("unexpected search result: %v %v", err, results)
	}
}

func TestService_Modify(t *testing.T) {
	store := NewMemoryStore()
	modStore := NewMemoryModStore()
	svc := &Service{Store: store, ModStore: modStore, LLM: FakeLLM{}}
	ctx := context.Background()
	r, err := svc.Create(ctx, 1, CreateRequest{Title: "Soup", Ingredients: []string{"carrot"}, Steps: []string{"mix"}})
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	newR, err := svc.Modify(ctx, 1, r.ID, "spicy")
	if err != nil {
		t.Fatalf("modify: %v", err)
	}
	if newR.ID == r.ID || !strings.Contains(newR.Title, "spicy") {
		t.Fatalf("unexpected modified recipe: %+v", newR)
	}
	if len(modStore.data) != 1 {
		t.Fatalf("mod not stored")
	}
}

func TestService_Modify_Err(t *testing.T) {
	store := NewMemoryStore()
	svc := &Service{Store: store, ModStore: NewMemoryModStore(), LLM: FakeLLM{}}
	ctx := context.Background()
	if _, err := svc.Modify(ctx, 1, 99, "ok"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected not found, got %v", err)
	}
	r, _ := svc.Create(ctx, 1, CreateRequest{Title: "Toast", Ingredients: []string{"bread"}, Steps: []string{"toast"}})
	if _, err := svc.Modify(ctx, 1, r.ID, "fail"); !errors.Is(err, ErrGenerationFailed) {
		t.Fatalf("expected generation failure")
	}
}
