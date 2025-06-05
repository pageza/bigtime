package user

import (
	"context"
	"testing"
)

func TestService_RegisterAndAuthenticate(t *testing.T) {
	store := NewMemoryStore()
	svc := &Service{Store: store, Secret: []byte("secret")}
	ctx := context.Background()

	u, err := svc.Register(ctx, "a@example.com", "pass")
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if u.ID == 0 {
		t.Fatalf("expected ID set")
	}

	token, err := svc.Authenticate(ctx, "a@example.com", "pass")
	if err != nil || token == "" {
		t.Fatalf("authenticate failed: %v token:%q", err, token)
	}

	if _, err := svc.Authenticate(ctx, "a@example.com", "bad"); err == nil {
		t.Fatal("expected bad password error")
	}
}
