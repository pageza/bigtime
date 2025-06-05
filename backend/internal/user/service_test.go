package user

import (
	"context"
	"testing"
)

func TestService_Register(t *testing.T) {
	store := NewMemoryStore()
	svc := &Service{Store: store}

	ctx := context.Background()
	u, err := svc.Register(ctx, "test@example.com", "secret")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if u.ID == 0 {
		t.Fatal("expected ID set")
	}
	if len(u.PasswordHash) == 0 {
		t.Fatal("expected password hash")
	}
	if _, err := svc.Register(ctx, "test@example.com", "another"); err == nil {
		t.Fatal("expected duplicate email error")
	}
}
