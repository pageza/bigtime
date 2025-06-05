package user

import "time"

// User represents a registered user.
type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}
