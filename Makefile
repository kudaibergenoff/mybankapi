# Проверка наличия файла .env
ifeq (,$(wildcard .env))
$(error .env file is missing)
endif

# Подключаем и экспортируем переменные из .env
include .env
export $(shell sed 's/=.*//' .env)

# Собираем строку подключения к базе данных
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

.PHONY: migrate migrate_down migrate_up migrate_version

# ==============================================================================
# Go migrate postgresql

force:
	migrate -database $(DATABASE_URL) -path migrations force 1

version:
	migrate -database $(DATABASE_URL) -path migrations version

migrate_up:
	migrate -database $(DATABASE_URL) -path migrations up 1

migrate_down:
	migrate -database $(DATABASE_URL) -path migrations down 1