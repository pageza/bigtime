package recipes

import (
	"context"

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
