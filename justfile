# Init
set dotenv-load

# Run the Gin server
[group('Server')]
run:
  @echo "Running Gin server...\n"
  go run main.go

# Run the Gin server in dev mode (with hot reload)
[group('Server')]
dev:
    @echo "Running gin server with hot reload..."
    air

### Database (host dev: db in docker, api on host)

# Start the database container and wait until it is ready
[group('Database')]
db-up:
    docker compose up -d db
    @echo "Waiting for database to be ready..."
    @until docker compose exec -T db pg_isready -U "$DB_USER" -d "$DB_NAME" > /dev/null; do sleep 1; done

# Stop the database container (data volume is kept)
[group('Database')]
db-down:
    docker compose stop db

[group('Database')]
db-logs:
    docker compose logs -f db

# Seed Pokemon data (only after fresh DB + migrations)
[group('Database')]
seed:
    @echo "Seeding database..."
    go run database/main.go

# One-time dev setup: start db, migrate, seed
[group('Database')]
db-init:
    just db-up
    just migrate-up
    just seed

### Migrations

# TODO: Make sslmode depend on env("ENV")
DB_URL := "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Apply migrations
[group('Migration')]
migrate-up:
    migrate -path database/migrations -database "{{DB_URL}}" up

# Rollback last migration
[group('Migration')]
migrate-down:
    migrate -path database/migrations -database "{{DB_URL}}" down 1

# Create new named migration
[group('Migration')]
migrate-create name:
    migrate create -ext sql -dir database/migrations {{name}}
