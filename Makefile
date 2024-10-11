ifeq (,$(wildcard .env))
$(error .env file is missing)
endif

include .env
export $(shell sed 's/=.*//' .env)

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

.PHONY: migrate migrate_down migrate_up migrate_version

# ==============================================================================
# Go migrate postgresql

force:
	echo "Stating migration force"
	migrate -database $(DATABASE_URL) -path migrations force 1

version:
	migrate -database $(DATABASE_URL) -path migrations version

migrate_up:
	echo "Starting migration up"
	migrate -database $(DATABASE_URL) -path migrations up 1

migrate_down:
	echo "Stating migration down"
	migrate -database $(DATABASE_URL) -path migrations down 1


# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker compose -f compose.yml up --build
down:
	echo "Down docker environment"
	docker compose -f compose.yml down
stop:
	echo "Stopping docker environment"
	docker compose -f compose.yml stop
build:
	echo "Building docker environment"
	docker compose -f compose.yml build
bash:
	docker compose -f compose.yml exec app /bin/sh

# ==============================================================================
# Main

run:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

doc-generate:
	swag init --dir cmd/ --generalInfo main.go