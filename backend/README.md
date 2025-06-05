# Backend

This directory contains the Go backend for **Alchemorsel**.

## Requirements

- Go 1.20+

## Running

```bash
go run ./cmd/api
```

The default server exposes a `/healthz` endpoint on port `8080`.

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
