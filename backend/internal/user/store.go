package user

import (
	"context"
	"errors"
	"sync"
)

// ErrNotFound indicates no user was found for the query.
var ErrNotFound = errors.New("user not found")

// Store provides persistence for users.
type Store interface {
	Create(ctx context.Context, u User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}

// MemoryStore is an in-memory Store implementation.
type MemoryStore struct {
	mu    sync.Mutex
	next  int64
	users map[string]User
}

// NewMemoryStore creates a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{users: make(map[string]User)}
}

// Create stores a new user.
func (m *MemoryStore) Create(ctx context.Context, u User) (User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.next++
	u.ID = m.next
	m.users[u.Email] = u
	return u, nil
}

// FindByEmail retrieves a user by email.
func (m *MemoryStore) FindByEmail(ctx context.Context, email string) (User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u, ok := m.users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
