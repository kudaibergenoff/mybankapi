ifeq (,$(wildcard .env))
$(error .env file is missing)
endif

include .env
export $(shell sed 's/=.*//' .env)

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

.PHONY: migrate migrate_down migrate_up migrate_version

# ==============================================================================
# Go migrate postgresql

force:
	migrate -database $(DATABASE_URL) -path migrations force

version:
	migrate -database $(DATABASE_URL) -path migrations version

migrate_up:
	migrate -database $(DATABASE_URL) -path migrations up

migrate_down:
	migrate -database $(DATABASE_URL) -path migrations down


# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f compose.yml up --build