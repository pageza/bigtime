# Backend

This directory contains the Go backend for **Alchemorsel**.

## Requirements
- Go 1.20+

## Running

```bash
go run ./cmd/api
```

The default server exposes a `/healthz` endpoint on port `8080`.

## Example

Below is a quick example showing how to use the `user.Service` in code:

```go
package main

import (
    "context"
    "fmt"

    "alchemorsel/internal/user"
)

func main() {
    store := user.NewMemoryStore()
    svc := &user.Service{Store: store, Secret: []byte("secret")}

    // Register a new user.
    u, err := svc.Register(context.Background(), "me@example.com", "pass")
    if err != nil {
        panic(err)
    }

    // Authenticate and get a signed token.
    token, err := svc.Authenticate(context.Background(), u.Email, "pass")
    if err != nil {
        panic(err)
    }
fmt.Println(token)
}
```

## Authentication Endpoints

Passwords are hashed using **Argon2** before being stored.

### Register

```bash
curl -X POST -d '{"email":"user@example.com","password":"password"}' \
  http://localhost:8080/v1/users
```

### Login

```bash
curl -X POST -d '{"email":"user@example.com","password":"password"}' \
  http://localhost:8080/v1/tokens
```

## Recipe Endpoints

### Create Recipe

```bash
curl -X POST -d '{"title":"Soup","ingredients":["water"],"steps":["boil"]}' \
  http://localhost:8080/v1/recipes
```

### Update Recipe

```bash
curl -X PUT -d '{"title":"Better Soup","ingredients":["water"],"steps":["boil"]}' \
  http://localhost:8080/v1/recipes/1
```

### Delete Recipe

```bash
curl -X DELETE http://localhost:8080/v1/recipes/1
```

### Modify Recipe

```bash
curl -X POST -d '{"prompt":"spicy"}' \
  http://localhost:8080/v1/recipes/1/modify
```

## Profile Endpoints

### Get Profile

```bash
curl http://localhost:8080/v1/profile
```

### Update Profile

```bash
curl -X PUT -d '{"displayName":"Bob"}' \
  http://localhost:8080/v1/profile
```
