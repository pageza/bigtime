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
