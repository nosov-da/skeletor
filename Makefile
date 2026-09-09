.RECIPEPREFIX = >

APP_NAME ?= api
BUILD_DIR ?= bin
DOCKER_IMAGE ?= $(APP_NAME)
DOCKER_TAG ?= latest
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/app?sslmode=disable

.PHONY: all run build test lint fmt vet tidy docker-build docker-up docker-down migrate-up migrate-down

all: fmt vet lint test build

run:
> go run ./cmd/api

build:
> go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api

test:
> go test -race -cover ./...

lint:
> golangci-lint run ./...

fmt:
> gofmt -s -w .

vet:
> go vet ./...

tidy:
> go mod tidy

docker-build:
> docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-up:
> docker compose up --build

docker-down:
> docker compose down

migrate-up:
> migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
> migrate -path migrations -database "$(DATABASE_URL)" down 1
