package recipes

import (
	"context"
	"errors"

	"strings"


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


// SearchRequest defines filters for querying recipes.
type SearchRequest struct {
	Q          string
	Ingredient []string
	Tag        []string
	Page       int
	Limit      int
}

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


// Search returns recipes matching the request filters.
func (s *Service) Search(ctx context.Context, req SearchRequest) ([]*Recipe, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 || req.Limit > 50 {
		req.Limit = 20
	}
	all, err := s.Store.List(ctx)
	if err != nil {
		return nil, err
	}
	var filtered []*Recipe
	for _, r := range all {
		if req.Q != "" && !strings.Contains(strings.ToLower(r.Title), strings.ToLower(req.Q)) {
			continue
		}
		match := true
		for _, ing := range req.Ingredient {
			found := false
			for _, rIng := range r.Ingredients {
				if strings.EqualFold(rIng, ing) {
					found = true
					break
				}
			}
			if !found {
				match = false
				break
			}
		}
		if !match {
			continue
		}
		for _, tag := range req.Tag {
			found := false
			for _, rTag := range r.Tags {
				if strings.EqualFold(rTag, tag) {
					found = true
					break
				}
			}
			if !found {
				match = false
				break
			}
		}
		if match {
			filtered = append(filtered, r)
		}
	}
	start := (req.Page - 1) * req.Limit
	if start >= len(filtered) {
		return []*Recipe{}, nil
	}
	end := start + req.Limit
	if end > len(filtered) {
		end = len(filtered)
	}
	return filtered[start:end], nil
}

