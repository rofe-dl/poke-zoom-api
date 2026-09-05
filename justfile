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

### Migrations

# TODO: Make sslmode depend on env("ENV")
DB_URL := "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST_WITHOUT_DOCKER}:${DB_PORT}/${DB_NAME}?sslmode=disable"

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