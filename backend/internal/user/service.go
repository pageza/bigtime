package user

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

// Service provides user management operations.
type Service struct {
	Store  Store
	Secret []byte
}

// Register creates a new user account.
func (s *Service) Register(ctx context.Context, email, password string) (User, error) {
	if _, err := s.Store.FindByEmail(ctx, email); err == nil {
		return User{}, errors.New("email already exists")
	} else if !errors.Is(err, ErrNotFound) {
		return User{}, err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return User{}, err
	}
	u := User{Email: email, PasswordHash: hash, CreatedAt: time.Now()}
	return s.Store.Create(ctx, u)
}

// Authenticate verifies credentials and returns a signed token.
func (s *Service) Authenticate(ctx context.Context, email, password string) (string, error) {
	u, err := s.Store.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if ok := verifyPassword(u.PasswordHash, password); !ok {
		return "", errors.New("invalid credentials")
	}
	return createToken(u.ID, s.Secret)
}

func hashPassword(pw string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(pw), salt, 1, 64*1024, 4, 32)
	return hex.EncodeToString(salt) + ":" + hex.EncodeToString(hash), nil
}

func verifyPassword(stored, pw string) bool {
	parts := strings.Split(stored, ":")
	if len(parts) != 2 {
		return false
	}
	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false
	}
	hash, err := hex.DecodeString(parts[1])
	if err != nil {
		return false
	}
	other := argon2.IDKey([]byte(pw), salt, 1, 64*1024, 4, 32)
	return subtle.ConstantTimeCompare(hash, other) == 1
}

func createToken(id int64, secret []byte) (string, error) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := fmt.Sprintf(`{"sub":%d,"exp":%d}`, id, time.Now().Add(time.Hour).Unix())
	payloadEnc := base64.RawURLEncoding.EncodeToString([]byte(payload))
	signing := header + "." + payloadEnc
	h := hmac.New(sha256.New, secret)
	if _, err := h.Write([]byte(signing)); err != nil {
		return "", err
	}
	sig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return signing + "." + sig, nil
}
