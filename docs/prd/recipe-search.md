# Recipe Search and Browse

## Summary

Provide users with robust search and browsing functionality so they can easily find recipes by keyword, ingredient, or tag. Supports pagination and sorting for large result sets.

## User Story

As a user, I want to search for recipes matching certain terms or ingredients so that I can quickly discover meals to cook.

## Requirements

- REST endpoint to search recipes with query parameters for text, ingredients, and tags.
- Results must be paginated and sortable by relevance or popularity.
- Errors for invalid queries or missing parameters must be descriptive.
- Search is available to both authenticated and anonymous users.

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

| Method | Path            | Description                  |
| ------ | --------------- | ---------------------------- |
| GET    | `/v1/recipes`   | Search and list recipes      |
| GET    | `/v1/recipes/:id` | Retrieve recipe details |

Query parameters for `/v1/recipes`:
- `q` search text
- `ingredient` filter by ingredient (repeatable)
- `tag` filter by tag (repeatable)
- `page` page number
- `limit` results per page

## Edge Cases

- No recipes match the query.
- Invalid page or limit values.
- Very long search terms.

## Acceptance Criteria

- Searching with no filters returns the first page of recipes.
- Filters by ingredient and tag narrow results correctly.
- Pagination works and returns total count.
- Unit tests cover no results, valid searches, and invalid parameter cases.
