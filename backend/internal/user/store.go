package user

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Predefined errors for user store operations.
var (
	ErrEmailExists = errors.New("email already exists")
	ErrNotFound    = errors.New("user not found")
)

// Store defines persistence methods for users.
type Store interface {
	Create(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	mu     sync.Mutex
	nextID int64
	users  map[string]*User
}

// NewMemoryStore initializes a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{users: make(map[string]*User)}
}

// Create inserts a new user if the email is not taken.
func (m *MemoryStore) Create(ctx context.Context, u *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.users[u.Email]; ok {
		return ErrEmailExists
	}
	m.nextID++
	u.ID = m.nextID
	u.CreatedAt = time.Now()
	m.users[u.Email] = u
	return nil
}

// GetByEmail returns a user by their email address.
func (m *MemoryStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u, ok := m.users[email]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}
