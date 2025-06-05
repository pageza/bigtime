package profile

import (
	"context"
	"testing"
)

func TestMemoryStore_CRUD(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()

	p := &Profile{UserID: 1, DisplayName: "A"}
	if err := store.Create(ctx, p); err != nil {
		t.Fatalf("create: %v", err)
	}
	if p.ID == 0 {
		t.Fatalf("id not set")
	}

	got, err := store.GetByUserID(ctx, 1)
	if err != nil || got.DisplayName != "A" {
		t.Fatalf("get: %v", err)
	}

	p.DisplayName = "B"
	if err := store.Update(ctx, p); err != nil {
		t.Fatalf("update: %v", err)
	}
	got, _ = store.GetByUserID(ctx, 1)
	if got.DisplayName != "B" {
		t.Fatalf("update not saved")
	}
}
