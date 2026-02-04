APP_NAME      ?= jwt-login
MAIN_PKG      ?= ./cmd

MIGRATIONS_DIR ?= migrations
MIGRATE_BIN    ?= migrate

DB_HOST     ?= localhost
DB_PORT     ?= 5432
DB_NAME     ?= taskmanager
DB_USER     ?= daf-devs
DB_PASSWORD ?= 1q2w3e4r
DB_SSLMODE  ?= disable

DB_URL ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)


ifneq (,$(wildcard .env))
include .env
export
endif

.PHONY: run
run: 
	go run $(MAIN_PKG)

.PHONY: build


.PHONY: tidy
tidy:
	go mod tidy

.PHONY: migrate-up
migrate-up: guard-migrate guard-migrations 
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down: guard-migrate guard-migrations 
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

.PHONY: migrate-down-all
migrate-down-all: guard-migrate guard-migrations 
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

.PHONY: migrate-force
migrate-force: guard-migrate guard-migrations 
	@test -n "$(V)" || (echo "Missing V. Example: make migrate-force V=1" && exit 1)
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(V)

.PHONY: migrate-version
migrate-version: guard-migrate guard-migrations
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

.PHONY: migrate-new
migrate-new: guard-migrate guard-migrations 
	@test -n "$(NAME)" || (echo "Missing NAME. Example: make migrate-new NAME=add_tasks_table" && exit 1)
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

