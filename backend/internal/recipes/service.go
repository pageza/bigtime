package recipes

import (
	"context"
	"errors"
)

// Service provides recipe creation operations.
type Service struct {
	Store Store
}

// CreateRequest defines input for creating a recipe.
type CreateRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
	Tags        []string `json:"tags"`
}

// ErrInvalidInput is returned when the request is missing required fields.
var ErrInvalidInput = errors.New("invalid recipe input")

// Create stores a new recipe owned by the given user.
func (s *Service) Create(ctx context.Context, userID int64, req CreateRequest) (*Recipe, error) {
	if req.Title == "" || len(req.Ingredients) == 0 || len(req.Steps) == 0 {
		return nil, ErrInvalidInput
	}
	r := &Recipe{
		Title:       req.Title,
		Description: req.Description,
		Ingredients: append([]string(nil), req.Ingredients...),
		Steps:       append([]string(nil), req.Steps...),
		Tags:        append([]string(nil), req.Tags...),
		CreatedBy:   userID,
	}
	if err := s.Store.Create(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}
