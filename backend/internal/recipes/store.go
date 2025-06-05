package recipes

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Store defines persistence operations for recipes.
type Store interface {
	Create(ctx context.Context, r *Recipe) error
	FindByID(ctx context.Context, id int64) (*Recipe, error)
	Update(ctx context.Context, r *Recipe) error
	Delete(ctx context.Context, id int64) error
}

// ErrNotFound is returned when a recipe cannot be found.
var ErrNotFound = errors.New("recipe not found")

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	mu     sync.RWMutex
	nextID int64
	data   map[int64]*Recipe
}

// NewMemoryStore returns an empty MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int64]*Recipe)}
}

// Create stores a new recipe and assigns an ID.
func (m *MemoryStore) Create(ctx context.Context, r *Recipe) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nextID++
	r.ID = m.nextID
	r.CreatedAt = time.Now()
	r.UpdatedAt = r.CreatedAt
	clone := *r
	m.data[r.ID] = &clone
	return nil
}

// FindByID retrieves a recipe by ID.
func (m *MemoryStore) FindByID(ctx context.Context, id int64) (*Recipe, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	clone := *r
	return &clone, nil
}

// Update modifies an existing recipe.
func (m *MemoryStore) Update(ctx context.Context, r *Recipe) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[r.ID]; !ok {
		return ErrNotFound
	}
	r.UpdatedAt = time.Now()
	clone := *r
	m.data[r.ID] = &clone
	return nil
}

// Delete removes a recipe by ID.
func (m *MemoryStore) Delete(ctx context.Context, id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[id]; !ok {
		return ErrNotFound
	}
	delete(m.data, id)
	return nil
}
