package users

import (
	"context"
	"errors"
	"sync"
	"time"
)

// User represents a registered account.
type User struct {
	ID           int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

// Store defines persistence behavior for Users.
type Store interface {
	Create(ctx context.Context, u *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

var ErrEmailExists = errors.New("email already in use")
var ErrNotFound = errors.New("user not found")

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	mu     sync.RWMutex
	users  map[string]*User
	nextID int64
}

// NewMemoryStore creates an empty MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{users: make(map[string]*User)}
}

// Create stores a new user if the email is unused.
func (m *MemoryStore) Create(ctx context.Context, u *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.users[u.Email]; ok {
		return ErrEmailExists
	}
	m.nextID++
	u.ID = m.nextID
	u.CreatedAt = time.Now()
	clone := *u
	m.users[u.Email] = &clone
	return nil
}

// FindByEmail returns a user by email.
func (m *MemoryStore) FindByEmail(ctx context.Context, email string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	u, ok := m.users[email]
	if !ok {
		return nil, ErrNotFound
	}
	clone := *u
	return &clone, nil
}
