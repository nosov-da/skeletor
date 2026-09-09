# Go Production Skeleton

Production-ready Go project skeleton.

## Quick start

1. Replace module path in go.mod and all Go imports:
   github.com/yourorg/yourproject

2. Download dependencies:
   go get github.com/go-chi/chi/v5@latest
   go get github.com/caarlos0/env/v11@latest
   go mod tidy

3. Copy environment example:
   cp configs/.env.example .env

4. Run locally:
   make run

## Docker

Run app and PostgreSQL:

    docker compose up --build

Health endpoints:

    GET /healthz
    GET /readyz

## Makefile

The Makefile uses GNU Make .RECIPEPREFIX.
If your make implementation does not support it, run commands manually.

Common commands:

    make run
    make build
    make test
    make lint
    make fmt
    make vet
    make tidy
    make docker-build
    make docker-up
    make docker-down
    make migrate-up
    make migrate-down

## Architecture

- internal/entity - domain models
- internal/repository - data access layer
- internal/service - business logic
- internal/handler - HTTP handlers
- internal/server - router and middleware
- internal/app - dependency wiring and application lifecycle
- internal/config - configuration
- internal/pkg/logger - structured logger

## Production checklist

Before production, add or configure:

- real database driver, for example pgx or sqlx
- migrations runner, for example golang-migrate
- authentication and authorization
- request validation
- metrics, for example Prometheus
- tracing, for example OpenTelemetry
- rate limiting
- API documentation, for example OpenAPI
- secrets management
