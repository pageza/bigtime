# Alchemorsel

Alchemorsel is an AI-powered recipe application. This repository contains both the backend (Go) and frontend (Vue 3) code.

## Structure

- `/backend` – Go services
- `/frontend` – Vue 3 app
- `/docs` – Documentation and PRDs

## Getting Started

```bash
# Backend
cd backend
go run ./cmd/api

# Frontend
cd frontend
yarn dev
```

## Testing

Run backend and frontend checks before committing changes:

```bash
cd backend
go test ./...
golangci-lint run ./...

cd ../frontend
npm run lint
```

See individual directories for details.
