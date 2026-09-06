# Poke-zoom API

The REST API for my game Poke-zoom. It uses Go with Huma, with Gin as the underlying router and Postgres for the database.

## Setup

### Docker

1. Set env variables in `.env` file. Make sure to set `DB_HOST=db`.
1. Run

   ```
   docker compose up
   ```

### Without Docker

1. Install Go version `1.26.1` and Postgres and create a database for this project.
1. Run
   ```
   go mod download
   ```
1. Set env variables in `.env` file. Make sure to set `DB_HOST=localhost`.
1. Install the command runner [just](https://github.com/casey/just#installation).
1. Install [air](https://github.com/air-verse/air#installation) for hot reloading.
1. Run

   ```
   just dev
   ```

## Database

If you wanna make migrations or if you setup the app **WITHOUT** using Docker:

1. Make sure you have Go and [just](https://github.com/casey/just#installation) installed.
1. Install [golang-migrate](https://github.com/golang-migrate/migrate).
1. Run `just migrate-up` to migrate the changes. Make sure `DB_HOST=localhost` in your `.env` file.

1. All the Pokemon data has to be seeded into the database from the JSON file after you set up the database and run the migrations. To do that, run:

   ```bash
   go run database/main.go
   ```

To migrate down if necessary, run `just migrate-down` to go down 1 migration.

To create a new migration, do `just migrate-create {name}`.
