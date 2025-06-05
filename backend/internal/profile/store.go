package profile

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Store defines persistence operations for profiles.
type Store interface {
	GetByUserID(ctx context.Context, userID int64) (*Profile, error)
	Create(ctx context.Context, p *Profile) error
	Update(ctx context.Context, p *Profile) error
}

// ErrNotFound is returned when a profile cannot be located.
var ErrNotFound = errors.New("profile not found")

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	mu     sync.RWMutex
	nextID int64
	data   map[int64]*Profile
}

// NewMemoryStore returns a ready MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int64]*Profile)}
}

// GetByUserID returns a profile for the user.
func (m *MemoryStore) GetByUserID(ctx context.Context, userID int64) (*Profile, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.data[userID]
	if !ok {
		return nil, ErrNotFound
	}
	clone := *p
	return &clone, nil
}

// Create stores a new profile.
func (m *MemoryStore) Create(ctx context.Context, p *Profile) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nextID++
	p.ID = m.nextID
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt
	clone := *p
	m.data[p.UserID] = &clone
	return nil
}

// Update modifies an existing profile.
func (m *MemoryStore) Update(ctx context.Context, p *Profile) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	existing, ok := m.data[p.UserID]
	if !ok {
		return ErrNotFound
	}
	p.ID = existing.ID
	p.CreatedAt = existing.CreatedAt
	p.UpdatedAt = time.Now()
	clone := *p
	m.data[p.UserID] = &clone
	return nil
}
