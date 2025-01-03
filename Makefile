APP_NAME = server

# GOOSE PARAMETERS
GOOSE_DRIVER=mysql
GOOSE_DBSTRING="root:root1234@tcp(localhost:3306)/shopdevgo"
GOOSE_MIGRATION_DIR=./sql/schema

dev:
	go run ./cmd/${APP_NAME}/main.go

goose-up:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) up

goose-down:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) down

goose-reset:
	goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate

swag:
	swag init -g ./cmd/${APP_NAME}/main.go -o ./docs
