package recipes

import (
	"context"
	"testing"
)

func TestMemoryStore_CRUD(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()
	r := &Recipe{Title: "Toast", Ingredients: []string{"Bread"}, Steps: []string{"Toast bread"}, CreatedBy: 1}
	if err := store.Create(ctx, r); err != nil {
		t.Fatalf("create: %v", err)
	}
	if r.ID == 0 {
		t.Fatalf("expected id set")
	}
	got, err := store.FindByID(ctx, r.ID)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got.Title != r.Title {
		t.Fatalf("title mismatch")
	}
	r.Title = "Toast 2"
	if err := store.Update(ctx, r); err != nil {
		t.Fatalf("update: %v", err)
	}
	got, _ = store.FindByID(ctx, r.ID)
	if got.Title != "Toast 2" {
		t.Fatalf("update not saved")
	}
	if err := store.Delete(ctx, r.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := store.FindByID(ctx, r.ID); err != ErrNotFound {
		t.Fatalf("expected not found after delete")
	}
}

func TestMemoryStore_List(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		r := &Recipe{Title: "T", Ingredients: []string{"i"}, Steps: []string{"s"}, CreatedBy: 1}
		if err := store.Create(ctx, r); err != nil {
			t.Fatalf("create: %v", err)
		}
	}
	list, err := store.List(ctx)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 3 {
		t.Fatalf("expected 3, got %d", len(list))
	}
	if list[0].ID > list[1].ID {
		t.Fatalf("list not sorted")
	}
}
