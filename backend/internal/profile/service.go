package profile

import (
	"context"
	"errors"
	"net/url"
)

// Service manages user profiles.
type Service struct {
	Store Store
}

// UpdateRequest defines editable profile fields.
type UpdateRequest struct {
	DisplayName        string   `json:"displayName"`
	AvatarURL          string   `json:"avatarURL"`
	Bio                string   `json:"bio"`
	DietaryPreferences []string `json:"dietaryPreferences"`
}

// ErrInvalidInput is returned for bad profile data.
var ErrInvalidInput = errors.New("invalid profile input")

// Get returns the user's profile.
func (s *Service) Get(ctx context.Context, userID int64) (*Profile, error) {
	return s.Store.GetByUserID(ctx, userID)
}

// Update creates or updates a profile.
func (s *Service) Update(ctx context.Context, userID int64, req UpdateRequest) (*Profile, error) {
	if req.DisplayName == "" || len(req.DisplayName) > 50 {
		return nil, ErrInvalidInput
	}
	if req.AvatarURL != "" {
		u, err := url.ParseRequestURI(req.AvatarURL)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
			return nil, ErrInvalidInput
		}
	}

	p, err := s.Store.GetByUserID(ctx, userID)
	if errors.Is(err, ErrNotFound) {
		p = &Profile{UserID: userID}
		p.DisplayName = req.DisplayName
		p.AvatarURL = req.AvatarURL
		p.Bio = req.Bio
		p.DietaryPreferences = append([]string(nil), req.DietaryPreferences...)
		if err := s.Store.Create(ctx, p); err != nil {
			return nil, err
		}
		return p, nil
	} else if err != nil {
		return nil, err
	}
	p.DisplayName = req.DisplayName
	p.AvatarURL = req.AvatarURL
	p.Bio = req.Bio
	p.DietaryPreferences = append([]string(nil), req.DietaryPreferences...)
	if err := s.Store.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
