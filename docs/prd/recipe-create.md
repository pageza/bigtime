# Recipe Creation and Editing

## Summary

Users can create new recipes manually or generate them via LLM assistance.

## User Story

As a registered user, I want to create and edit recipes so that I can keep my cooking ideas organized.

## Requirements

- CRUD endpoints for recipes owned by a user.
- Optional LLM-powered generation endpoint that returns a new recipe based on user prompts.
- Only authenticated users may create, edit, or delete recipes.
- Errors must be descriptive and actionable.

## Data Model

```
Recipe {
  ID int64
  UserID int64
  Title string
  Ingredients []string
  Steps []string
  CookTimeMinutes int
  CreatedAt time.Time
}
```

## API Endpoints

| Method | Path                     | Description                 |
| ------ | ------------------------ | --------------------------- |
| POST   | `/v1/recipes`            | Create new recipe           |
| GET    | `/v1/recipes/{id}`       | Retrieve recipe             |
| PUT    | `/v1/recipes/{id}`       | Update recipe               |
| DELETE | `/v1/recipes/{id}`       | Delete recipe               |
| POST   | `/v1/recipes/generate`   | Generate recipe via LLM     |

## Edge Cases

- Invalid ingredient list
- Unauthorized update/delete
- LLM generation failure

## Acceptance Criteria

- CRUD operations return appropriate status codes.
- LLM generation returns 201 with generated recipe data.
- Unit tests cover success and failure scenarios.
