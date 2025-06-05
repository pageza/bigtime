package profile

import "time"

// Profile holds user-specific information.
type Profile struct {
	ID                 int64
	UserID             int64
	DisplayName        string
	AvatarURL          string
	Bio                string
	DietaryPreferences []string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
