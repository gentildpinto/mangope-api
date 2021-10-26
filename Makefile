DATABASE_HOST     ?= $(shell grep "DATABASE_HOST" .env | cut -d '=' -f2)
DATABASE_PORT     ?= $(shell grep "DATABASE_PORT" .env | cut -d '=' -f2)
DATABASE_NAME 	  ?= $(shell grep "DATABASE_NAME" .env | cut -d '=' -f2)
DATABASE_USERNAME ?= $(shell grep "DATABASE_USERNAME" .env | cut -d '=' -f2)
DATABASE_PASSWORD ?= $(shell grep "DATABASE_PASSWORD" .env | cut -d '=' -f2)
# postgres://user:password@host:port/dbname?query
PGSQL_DSN ?= ${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}

# Version of migrations - this is optionally used on goto command
V?=

# Number of migrations - this is optionally used on up and down commands
N?=

migrate-setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing golang-migrate..."; go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; fi

.PHONY: migrate-up
migrate-up: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations up $(N)

.PHONY: migrate-down
migrate-down: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations down $(N)

.PHONY: migrate-to-version
migrate-to-version: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations goto $(V)

.PHONY: drop-db
drop-db: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations drop

.PHONY: force-version
force-version: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations force $(V)

.PHONY: migration-version
migration-version: migrate-setup
	@ migrate -database 'postgres://$(PGSQL_DSN)?sslmode=disable' -path $$(pwd)/database/migrations version

.PHONY: start-dev
start-dev:
	docker compose up