# Poke-zoom API

The REST API for my game Poke-zoom. It uses Go with Huma, with Gin as the underlying router and Postgres for the database.

## Setup

### Docker (full stack)

For trying out the API. Starts the database, runs migrations + seeding once, and runs the API.

1. Set env variables in `.env` file.
1. Run

   ```
   docker compose up
   ```

### Development (host API + Docker database)

For actual development work. The database runs in a Docker container, while the API, migrations, and seeding run on your machine.

1. Install Go version `1.26.1` and Docker.
1. Install the command runner [just](https://github.com/casey/just#installation).
1. Install [air](https://github.com/air-verse/air#installation) for hot reloading and [golang-migrate](https://github.com/golang-migrate/migrate).
1. Run
   ```
   go mod download
   ```
1. Set env variables in `.env` file. Make sure to set `DB_HOST=localhost`.
1. One-time setup: `just db-init` (starts the db container, runs migrations, seeds Pokemon data).
1. Run `just dev` (hot reload) or `just run`.

Useful commands: `just db-up`, `just db-down`, `just db-logs`, `just seed`.

## Database

Migrations and seeding run on your machine (or in the `migrate` container in the full Docker setup):

1. Run `just migrate-up` to apply the migrations. Make sure `DB_HOST` in your `.env` file points at your database (`localhost` for host dev, overridden to `db` inside Docker).

1. All the Pokemon data has to be seeded into the database from the JSON file after you set up the database and run the migrations. To do that, run:

   ```bash
   just seed
   ```
   If you are running the project from inside Docker, then this has already been handled for you.

To migrate down if necessary, run `just migrate-down` to go down 1 migration.

To create a new migration, do `just migrate-create {name}`.
