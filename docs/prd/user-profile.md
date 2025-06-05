# User Profiles

## Summary

Allow each user to view and update personal details such as display name, avatar, and dietary preferences. Profiles personalize recipe results and are required for other features like favorites and pantry management.

## User Story

As an authenticated user, I want to manage my profile so that Alchemorsel can tailor recipes and settings to my needs.

## Requirements

- REST endpoints to retrieve and update the current user's profile.
- Only authenticated users can modify their own profile.
- Profile includes display name, avatar URL, bio, and dietary preference fields.
- Validation for field lengths and required information.
- Descriptive error messages for invalid input or unauthorized access.

## Data Model

```
Profile {
  ID int64
  UserID int64
  DisplayName string
  AvatarURL string
  Bio string
  DietaryPreferences []string
  CreatedAt time.Time
  UpdatedAt time.Time
}
```

## API Endpoints

| Method | Path            | Description                |
| ------ | --------------- | -------------------------- |
| GET    | `/v1/profile`   | Get the authenticated user's profile |
| PUT    | `/v1/profile`   | Update the authenticated user's profile |

## Edge Cases

- Attempt to update another user's profile.
- Missing or invalid avatar URL.
- Display name exceeding length limits.

## Acceptance Criteria

- `GET /v1/profile` returns 200 with profile data for authenticated user.
- `PUT /v1/profile` returns 200 and updates the profile when data is valid.
- Validation errors return 400 with actionable messages.
- Unit tests cover success, validation errors, and unauthorized cases.
