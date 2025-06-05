# Backend

This directory contains the Go backend for **Alchemorsel**.

## Requirements
- Go 1.20+

## Running

```bash
go run ./cmd/api
```

The default server exposes a `/healthz` endpoint and a `/v1/users` registration API on port `8080`.
