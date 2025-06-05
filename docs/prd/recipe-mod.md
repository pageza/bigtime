# Recipe Modification Requests

## Summary

Allow users to submit modification requests for existing recipes. Requests are processed by the LLM to adjust ingredients, steps, or servings based on user preferences.

## User Story

As a user, I want to tweak existing recipes to suit my dietary needs or portion sizes so that I can cook meals tailored to me.

## Requirements

- REST endpoint to submit a modification request referencing an existing recipe ID and desired changes.
- LLM processes the request and returns an updated recipe version.
- Original recipe must remain unchanged; modifications are stored as new recipes linked to the source.
- Validation for reasonable serving sizes and ingredient substitutions.
- Clear error handling if the LLM fails to generate a valid modification.

## Data Model

```
RecipeModification {
  ID int64
  SourceRecipeID int64
  ModifiedRecipeID int64
  RequestedBy int64
  Prompt string
  CreatedAt time.Time
}
```

## API Endpoints

| Method | Path                       | Description                           |
| ------ | -------------------------- | ------------------------------------- |
| POST   | `/v1/recipes/:id/modify`   | Request a modified version of a recipe |

## Edge Cases

- Invalid recipe ID supplied.
- LLM unable to create a coherent modification.
- User submits conflicting or impossible requests.

## Acceptance Criteria

- Providing a valid prompt produces a new recipe based on the original.
- Original recipe is preserved and modification references it.
- Errors are returned for invalid requests or generation failures.
- Unit tests cover successful modifications, invalid IDs, and LLM failure cases.
