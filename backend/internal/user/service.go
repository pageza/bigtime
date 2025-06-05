package user

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// Service provides user-related business logic.
type Service struct {
	Store Store
}

// Register creates a new user with the provided credentials.
func (s *Service) Register(ctx context.Context, email, password string) (*User, error) {
	hash, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	u := &User{Email: email, PasswordHash: hash}
	if err := s.Store.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}

// hashPassword hashes the given password using SHA-256 with a random salt.
// This is a temporary stand-in for Argon2 which requires external
// dependencies not available in this environment.
func hashPassword(password string) ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	h := sha256.Sum256(append(salt, []byte(password)...))
	out := append(salt, h[:]...)
	return out, nil
}
