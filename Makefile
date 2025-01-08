APP_NAME = server

# GOOSE PARAMETERS
GOOSE_DRIVER=mysql
GOOSE_DBSTRING="root:root1234@tcp(localhost:3306)/shopdevgo"
GOOSE_MIGRATION_DIR=./sql/schema

NAME_SQL ?= abc

dev:
	go run ./cmd/${APP_NAME}/main.go

goose-create:
	goose -s create $(NAME_SQL) sql -dir=$(GOOSE_MIGRATION_DIR)

goose-up:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) up

# Migrate up a single migration from the current version
goose-up-by-one:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) up-by-one

goose-down:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) down

goose-reset:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate

swag:
	swag init -g ./cmd/${APP_NAME}/main.go -o ./docs
