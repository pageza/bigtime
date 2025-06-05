package recipes

import (
	"context"
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
