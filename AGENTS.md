# AGENTS.md – Alchemorsel Agent Guardrails

## Coding Standards
- Idiomatic Go for backend, idiomatic Vue 3 + Tailwind for frontend
- All code must be formatted using `gofmt` (Go) and Prettier (Vue/JS)
- Use context.Context for all Go handlers/services

## Project Structure
- Monorepo with `/backend` and `/frontend` at root
- `backend/cmd/api/main.go` as entrypoint
- Follow blueprint's folder/component structure exactly

## Workflow
1. For each feature, generate a PRD per the blueprint
2. Break PRD into atomic, testable tasks
3. For each task:
   - Write code
   - Write unit/integration test
   - Write inline doc comments
   - Write/extend README if applicable
   - Run tests and ensure 100% pass before proceeding

## Documentation
- Every exported function/type must have a doc comment
- Each component/service must include usage examples in README
- All new API endpoints must be documented in OpenAPI/Swagger

## Error Handling
- No panics in production logic
- All errors must be wrapped with context
- User-facing errors must be actionable and friendly

## Tests
- All Go code must have table-driven tests where feasible
- All Vue components must have at least shallow rendering/unit test

## Lint/Checks
- Run `golangci-lint run ./...` on backend before completion
- Run `npm run lint` on frontend

## Ambiguity
- If logic is missing or ambiguous, HALT and request clarification from user
- Never invent business logic without a directive

## Self-Check
- After every task, verify:
   - Code matches blueprint/PRD
   - All tests pass
   - Documentation updated
- At end of feature, summarize changes and any TODOs/risks

# End of AGENTS.md
