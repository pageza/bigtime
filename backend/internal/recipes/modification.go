package recipes

import (
	"context"
	"sync"
	"time"
)

// Modification links a generated recipe back to its source.
type Modification struct {
	ID               int64
	SourceRecipeID   int64
	ModifiedRecipeID int64
	RequestedBy      int64
	Prompt           string
	CreatedAt        time.Time
}

// ModStore defines persistence for recipe modifications.
type ModStore interface {
	Create(ctx context.Context, m *Modification) error
}

// MemoryModStore is an in-memory ModStore implementation.
type MemoryModStore struct {
	mu     sync.Mutex
	nextID int64
	data   map[int64]*Modification
}

// NewMemoryModStore returns an empty MemoryModStore.
func NewMemoryModStore() *MemoryModStore {
	return &MemoryModStore{data: make(map[int64]*Modification)}
}

// Create saves a modification entry.
func (m *MemoryModStore) Create(ctx context.Context, mod *Modification) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nextID++
	mod.ID = m.nextID
	mod.CreatedAt = time.Now()
	clone := *mod
	m.data[mod.ID] = &clone
	return nil
}
