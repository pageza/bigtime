# Alchemorsel v2 – Architecture Blueprint (SUPERCHARGED FOR AGENTS)

## 0.0. Purpose
This blueprint is designed for autonomous coding agents. Every requirement is explicit, testable, and structured to prevent ambiguity or agent drift. Agents MUST follow every directive, ask for clarification if logic is missing, and halt on unhandled ambiguity.

# 1. Project Overview
- App Name: Alchemorsel
- One-Sentence Description: A Recipe generating app, powered by AI.
- Vision Statement: An app to help people cut down on food waste, eat better and teach how to cook.
- Core Value Proposition: Modifies recipes for the user maintaining the integrity of the recipes. Helps user to tailor recipes and diet plans to their specs.
- Primary Users / Personas: See Appendice: User Persona Table
- Key Differentiators:
    - Does the ‘heavy-lifting’ of modifying or creating recipes that make sense (not substituting salt for sugar, etc.)
    - Interactive recipe design

# 2. User Stories & Personas
- Core User Stories, Edge Cases, Non-Goals, see original blueprint for expanded descriptions.
- Every user story must have a corresponding feature, acceptance test, and PRD.

# 3. Features (MVP + Planned)
**ALL features listed below MUST result in one PRD and an atomic implementation.**
- User registration/authentication
- User profiles
- Recipe search/browse
- Recipe creation/editing (manual + LLM-generated)
- Recipe modification by user request (LLM-driven)
- (See future features for full list, do not omit any planned features.)

# 4. App Structure
## Views/Screens, Component Inventory
- Follow the list from the original blueprint. Each view/component is an atomic deliverable. No ambiguity in naming or boundaries.

## Navigation/App Map
- Home, Search, Recipe Detail, Pantry, Profile, etc.

## State Management
- User Session: Global
- Recipes: Global
- Pantry Items: Global
- UI State: Local

# 5. Logic & Data Flow
## Authentication, Recipe Search, Generation, Creation, Detail, Profile, Pantry
- ALL app flows must have clear error/edge case handling.
- Every API call, DB query, LLM call, and UI event should be described, implemented, and tested.

## Data Models
- User, Recipe, PantryItem, Favorite
- Use explicit types, primary keys, relations, and indexes.
- Models MUST match both frontend and backend (no drift).

# 6. AI/LLM Integration
- RAG flows, LLM agent logic, MCP tools (see original)
- Agents must halt and request clarification for ambiguous flows/queries.

# 7. Tech Stack
- Backend: Go, REST, gRPC, GORM/sqlx, JWT, bcrypt/argon2, Swagger
- Frontend: Vue 3, Tailwind, Pinia, Vite, esbuild
- AI/LLM: DeepSeek, Ollama, OpenAI ADA
- Infra: AWS EC2, Docker Compose, RDS, S3
- CI/CD: GitHub Actions, Docker Compose, Automated tests
- Monitoring: Sentry, Prometheus/Grafana

# 8. Security, Privacy, and Compliance
- Explicit requirements for all listed compliance/standards (see original).
- Agents must not invent or relax any security logic.

# 9. Deployment Strategy
- Environment requirements, versioned DB migrations, rollback/runbook.
- Tests required for each stage.

# 10. Appendices
- Glossary, References, Open Questions/Risks, User Persona Table

# 11. Local Development Environment Setup
- Prerequisites, project scaffolding, Docker Compose, env/secrets, migrations, testing, endpoints, troubleshooting.

# 12. Agent-Specific Directives (ADDITIONAL FOR THIS VERSION)
- For each feature and component:
    - Generate a detailed PRD per MDC pipeline
    - Decompose PRD into atomic, testable tasks (no multi-part tasks)
    - For each task: write code, write a test, write doc/comments, run all tests
    - Do not proceed to next task until current is 100% passing
    - Summarize changes, outstanding TODOs, risks, and areas needing human review at the end of each feature
    - Use AGENTS.md for ALL style/logic/structure guidance. Never deviate unless explicitly told.

# 13. Deliverables (for the Agent)
- A `/backend` directory (Go)
- A `/frontend` directory (Vue 3)
- All atomic components as described
- A complete README, setup script, and project map
- All endpoints, models, migrations, seeds, tests, Swagger/OpenAPI docs
- All code committed and organized by feature/task
- One PRD per feature, one task list per PRD, one commit per task

---

Refer to the original "Alchemorsel-v2-Architecture-Blueprint-NORMALIZED.md" for all context, user stories, full edge cases, and non-goals.
