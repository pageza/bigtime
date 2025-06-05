package user

import "time"

// User represents a registered application user.
type User struct {
	ID           int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
