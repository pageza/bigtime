# Recipe Creation and Editing

## Summary

Enable users to create new recipes manually or generate them using the integrated LLM. Users can edit recipes they own at any time.

## User Story

As a cook, I want to create or generate recipes and edit them later so I can share my unique dishes with the community.

## Requirements

- REST endpoints to create, read, update, and delete recipes owned by the user.
- Optional flag to request AI generation of a recipe based on a prompt.
- Validation for required fields such as title, ingredients, and steps.
- Only the recipe owner can update or delete a recipe.
- Clear errors for invalid input or unauthorized actions.

## Data Model

```
Recipe {
  ID int64
  Title string
  Description string
  Ingredients []string
  Steps []string
  Tags []string
  CreatedBy int64
  CreatedAt time.Time
  UpdatedAt time.Time
}
```

## API Endpoints

| Method | Path               | Description                        |
| ------ | ------------------ | ---------------------------------- |
| POST   | `/v1/recipes`      | Create a new recipe (manual or AI) |
| PUT    | `/v1/recipes/:id`  | Update an existing recipe          |
| DELETE | `/v1/recipes/:id`  | Delete an existing recipe          |

Request body for AI generation includes a `prompt` field describing the desired recipe.

## Edge Cases

- Missing required fields when creating a recipe.
- Attempt to edit or delete a recipe not owned by the user.
- LLM generation fails or returns invalid output.

## Acceptance Criteria

- Recipes can be created manually with valid data.
- AI generation returns a new recipe when a prompt is provided.
- Only the owner can modify or delete their recipes.
- Unit tests cover manual creation, AI generation, update, and failure scenarios.
