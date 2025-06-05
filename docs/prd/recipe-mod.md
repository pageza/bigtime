# Recipe Modification via AI

## Summary

Allow users to modify existing recipes with AI assistance, such as adjusting servings or swapping ingredients.

## User Story

As a user, I want to tweak recipes with AI so that they better match my dietary needs or serving size.

## Requirements

- Endpoint accepts recipe ID and modification instructions.
- Uses LLM to generate a modified recipe variant.
- Original recipe must remain unchanged; modifications are stored as new recipe records.
- Handle LLM failures gracefully.

## Data Model

```
RecipeModification {
  ID int64
  OriginalRecipeID int64
  UserID int64
  ModificationRequest string
  CreatedAt time.Time
}
```

## API Endpoints

| Method | Path                              | Description                     |
| ------ | --------------------------------- | ------------------------------- |
| POST   | `/v1/recipes/{id}/modify`         | Modify recipe via AI            |

## Edge Cases

- Invalid modification request
- LLM service unavailable
- User tries to modify recipe they don't own

## Acceptance Criteria

- Successful modification returns 201 with new recipe ID.
- Failure scenarios return appropriate errors.
- Unit tests cover success and failure scenarios.
