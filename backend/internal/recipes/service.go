package recipes

import (
	"context"
	"errors"
	"strings"

	"alchemorsel/internal/profile"
)

// Service provides recipe creation operations.
type Service struct {
	Store        Store
	ModStore     ModStore
	LLM          LLM
	ProfileStore profile.Store
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

// PersonalizedSearch filters search results using the user's profile dietary preferences.
func (s *Service) PersonalizedSearch(ctx context.Context, userID int64, req SearchRequest) ([]*Recipe, error) {
	results, err := s.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	if s.ProfileStore == nil {
		return results, nil
	}
	p, err := s.ProfileStore.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, profile.ErrNotFound) {
			return results, nil
		}
		return nil, err
	}
	var filtered []*Recipe
	for _, r := range results {
		skip := false
		for _, ing := range r.Ingredients {
			for _, avoid := range p.DietaryPreferences {
				if strings.EqualFold(ing, avoid) {
					skip = true
					break
				}
			}
			if skip {
				break
			}
		}
		if !skip {
			filtered = append(filtered, r)
		}
	}
	return filtered, nil
}

// ErrInvalidPrompt is returned when the prompt is empty.
var ErrInvalidPrompt = errors.New("invalid modification prompt")

// ErrNotOwner is returned when a user attempts to modify another user's recipe.
var ErrNotOwner = errors.New("not recipe owner")

// Modify creates a modified recipe using the LLM and records the modification.
func (s *Service) Modify(ctx context.Context, userID, recipeID int64, prompt string) (*Recipe, error) {
	if prompt == "" {
		return nil, ErrInvalidPrompt
	}
	orig, err := s.Store.FindByID(ctx, recipeID)
	if err != nil {
		return nil, err
	}
	newRec, err := s.LLM.ModifyRecipe(ctx, orig, prompt)
	if err != nil {
		return nil, err
	}
	if err := s.Store.Create(ctx, newRec); err != nil {
		return nil, err
	}
	if s.ModStore != nil {
		_ = s.ModStore.Create(ctx, &Modification{
			SourceRecipeID:   orig.ID,
			ModifiedRecipeID: newRec.ID,
			RequestedBy:      userID,
			Prompt:           prompt,
		})
	}
	return newRec, nil
}

// Update modifies an existing recipe owned by the user.
func (s *Service) Update(ctx context.Context, userID, recipeID int64, req CreateRequest) (*Recipe, error) {
	if req.Title == "" || len(req.Ingredients) == 0 || len(req.Steps) == 0 {
		return nil, ErrInvalidInput
	}
	r, err := s.Store.FindByID(ctx, recipeID)
	if err != nil {
		return nil, err
	}
	if r.CreatedBy != userID {
		return nil, ErrNotOwner
	}
	r.Title = req.Title
	r.Description = req.Description
	r.Ingredients = append([]string(nil), req.Ingredients...)
	r.Steps = append([]string(nil), req.Steps...)
	r.Tags = append([]string(nil), req.Tags...)
	if err := s.Store.Update(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Delete removes a recipe owned by the user.
func (s *Service) Delete(ctx context.Context, userID, recipeID int64) error {
	r, err := s.Store.FindByID(ctx, recipeID)
	if err != nil {
		return err
	}
	if r.CreatedBy != userID {
		return ErrNotOwner
	}
	return s.Store.Delete(ctx, recipeID)
}
