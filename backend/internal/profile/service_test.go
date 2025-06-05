package profile

import (
	"context"
	"testing"
)

func TestService_Update(t *testing.T) {
	svc := &Service{Store: NewMemoryStore()}
	ctx := context.Background()

	p, err := svc.Update(ctx, 1, UpdateRequest{DisplayName: "Test"})
	if err != nil || p.DisplayName != "Test" {
		t.Fatalf("update: %v", err)
	}

	if _, err := svc.Update(ctx, 1, UpdateRequest{DisplayName: ""}); err == nil {
		t.Fatalf("expected validation error")
	}

	if _, err := svc.Update(ctx, 1, UpdateRequest{DisplayName: "Test", AvatarURL: "bad://"}); err == nil {
		t.Fatalf("expected url error")
	}
}
