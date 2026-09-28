# Agent Instructions

REST API for the game Poke-zoom. Go + [Huma](https://github.com/danielgtaylor/huma) (OpenAPI) on top of Gin, GORM, Postgres.

## Commands

- `just run` / `just dev` (dev = air hot reload, needs `.env`)
- `just migrate-up` / `just migrate-down` (rollback one)
- `just migrate-create <name>` (new SQL migration in `database/migrations`)
- Seed Pokemon data (only after fresh DB + migrations): `go run database/main.go`

There are no Go tests yet; verify changes by running the server and hitting endpoints.

## Layout

- `main.go` - loads `.env`, connects DB, wires Huma API under `/api/v1`
- `routes/` - Huma operation registration only; keep handlers thin
- `services/` - business logic and GORM queries (one service struct per domain, holds `*gorm.DB`)
- `models/` - domain models + request/response/query-param types
- `utils/` - error types and HTTP error mapping
- `database/` - migrations (`migrate` CLI, SQL), seeds, one-off scripts
- `tmp/` - scratch; don't add code here

## Conventions

- New endpoints: add a `register<X>Routes(api, db)` in `routes/` and call it from `routes.RegisterAll`. Use `huma.Register` with an `OperationID`, `Tags`, and a description so OpenAPI docs stay complete.
- Errors: return domain errors from services (e.g. `utils.NotFoundError`); routes map them with `utils.HandleHttpError` (unknown error -> 500). Don't return raw GORM errors to Huma.
- DB access only via service structs; no GORM calls in routes.
- Migrations are plain SQL, must be paired with a matching `.down.sql`.
- Env vars come from `.env` (see `sample.env`); `sslmode` is derived from `ENV` in `db.go`, don't hardcode it.

## Gotchas

- Never commit `.env` or secrets.
- `services/pokemon.go` contains a `time.Sleep(3s)` left in for testing - don't replicate that pattern.
- Pokemon data is seeded from a JSON file, not from migrations.
