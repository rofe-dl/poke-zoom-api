# Poke-zoom API

The REST API for my game Poke-zoom. It uses Go with Huma, with Gin as the underlying router and Postgres for the database.

## Setup

1. Install Go version `1.26.1`.
1. Set env variables in `.env` file.
1. Install the command runner [just](https://github.com/casey/just#installation).
1. Run

   ```
   docker compose up
   ```

## Database

### Migrations

Install [golang-migrate](https://github.com/golang-migrate/migrate), then the following commands will work:

1. `just migrate-up` to migrate up
2. `just migrate-down` to migrate down
3. `just migrate-create {name}` to create a migration

### Seed Data

All the Pokemon data has to be seeded into the database from the JSON file after you set up the database and run the migrations.

```bash
go run database/main.go
```
