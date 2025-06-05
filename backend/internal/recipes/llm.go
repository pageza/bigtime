package recipes

import (
	"context"
	"errors"
	"strings"
)

// LLM defines recipe modification generation behavior.
type LLM interface {
	ModifyRecipe(ctx context.Context, original *Recipe, prompt string) (*Recipe, error)
}

// ErrGenerationFailed is returned when the LLM cannot produce a result.
var ErrGenerationFailed = errors.New("llm generation failed")

// FakeLLM is a simple in-process implementation used for tests.
type FakeLLM struct{}

// ModifyRecipe returns a modified copy of the recipe unless the prompt contains "fail".
func (FakeLLM) ModifyRecipe(ctx context.Context, original *Recipe, prompt string) (*Recipe, error) {
	if strings.Contains(prompt, "fail") {
		return nil, ErrGenerationFailed
	}
	r := *original
	r.ID = 0
	r.Title = original.Title + " - " + prompt
	return &r, nil
}
