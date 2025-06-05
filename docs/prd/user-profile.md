# User Profiles

## Summary

Allow users to view and edit their profile details.

## User Story

As a logged in user, I want to manage my profile so that other users see accurate information about me.

## Requirements

- Endpoints to fetch and update a user's profile.
- Profile includes display name, bio, and avatar URL.
- Only authenticated users may update their own profile.
- Errors must be actionable and friendly.

## Data Model

```
Profile {
  ID int64
  UserID int64
  DisplayName string
  Bio string
  AvatarURL string
  CreatedAt time.Time
}
```

## API Endpoints

| Method | Path            | Description           |
| ------ | --------------- | --------------------- |
| GET    | `/v1/profile`   | Get current profile   |
| PUT    | `/v1/profile`   | Update current profile|

## Edge Cases

- Missing required fields on update
- Unauthorized update attempt
- Invalid avatar URL

## Acceptance Criteria

- Fetching profile returns 200 with profile data.
- Updating profile returns 200 when data valid.
- Unit tests cover success and failure scenarios.
