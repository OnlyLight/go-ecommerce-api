APP_NAME = server

# GOOSE PARAMETERS
GOOSE_DRIVER=mysql
GOOSE_DBSTRING="root:root1234@tcp(localhost:3306)/shopdevgo"
GOOSE_MIGRATION_DIR=./migrations/sql

run:
	go run ./cmd/${APP_NAME}/main.go

goose-up:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) up

goose-down:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) down

goose-reset:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) reset

sqlc-generate:
	sqlc generate

.PHONY: run
