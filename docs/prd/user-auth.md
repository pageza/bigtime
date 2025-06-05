# User Registration and Authentication

## Summary

Enable users to register and log in to Alchemorsel. Provides JWT-based authentication for protected endpoints.

## User Story

As a new visitor, I want to create an account and securely log in so that I can save and manage my recipes.

## Requirements

- REST endpoints for user registration and login.
- Passwords stored using Argon2.
- JWT tokens signed with server secret.
- Duplicate email addresses must be rejected.
- Errors should be descriptive and actionable.

## Data Model

```
User {
  ID int64
  Email string
  PasswordHash string
  CreatedAt time.Time
}
```

## API Endpoints

| Method | Path         | Description          |
| ------ | ------------ | -------------------- |
| POST   | `/v1/users`  | Register a new user  |
| POST   | `/v1/tokens` | Login and obtain JWT |

## Edge Cases

- Weak password
- Email already in use
- Invalid credentials during login

## Acceptance Criteria

- Registration returns 201 with JWT on success.
- Login returns 200 with JWT when credentials valid.
- Unit tests cover success and failure scenarios.
