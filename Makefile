.SILENT:

# Configuration
BINARY_DIR := ./
BINARY_NAME := cushon
MAIN_SRC := cmd/server/main.go
DB_MIGRATIONS_DIR := database/migrations
DB_LOCAL_URL := mysql://root:localdev@tcp(localhost:3306)/cushondb
DB_LOCAL_COMPOSE := .devcontainer/docker-compose.yml
VERSION_PKG := cushon/cmd/version

# Help target
.PHONY: help
help:
	echo "Usage:"
	echo "  make gen-api-v1      - Generate API v1 code"
	echo "  make gen-docs        - Generate API documentation in HTML format"
	echo "  make gen-models	     - Generate DTOs (models) from the genesis database schema"
	echo "  make build           - Compile the Go project and output the binary"
	echo "  make build-min       - Compile the Go project with -ldflags=\"-s -w\" and output the binary"
	echo "  make clean           - Remove the compiled binary"
	echo "  make start           - Start (run) the already compiled binary (doesn't build or migrate)"
	echo "  make test-unit       - Run unit tests"
	echo "  make test-cover      - Run unit tests with coverage"
	echo "  make test-int        - Run the integration tests"
	echo "  make update          - Update all packages in all subdirectories, tidy and verify"
	echo "  make local-db-up     - Start local database using docker-compose"
	echo "  make local-db-dn     - Stop local database using docker-compose"
	echo "  make migrate-install - Install migrate CLI"
	echo "  make migrate-create  - Create a new migration file"
	echo "  make migrate-up      - Run migrations"
	echo "  make migrate-down    - Rollback migrations"
	echo "  make run             - Build, start database, run migrations and then start the app"

.PHONY: gen-api-v1
gen-api-v1:
	oapi-codegen -package=api -generate types,fiber api/v1/openapi.yaml > internal/router/api/openapi.v1.gen.go

.PHONY: gen-docs-html
gen-doc:
	@if ! command -v redoc-cli >/dev/null 2>&1; then \
		echo "redoc-cli not found. Please install it with: npm install -g redoc-cli"; \
		exit 1; \
	fi
	redoc-cli bundle ./api/v1/openapi.yaml --output ./docs/api/service.html

.PHONY: gen-docs-md
gen-docs-md:
	@if ! command -v widdershins >/dev/null 2>&1; then \
		echo "widdershins not found. Please install it with: npm install -g widdershins"; \
		exit 1; \
	fi
	widdershins ./api/v1/openapi.yaml -o ./docs/api/service.md

# Helper to generate a self-signed certificate
.PHONY: gen-cert
gen-cert:
	openssl req -x509 -newkey rsa:2048 -keyout cert.key -out cert.crt -days 365 -nodes

.PHONY: gen-docs
gen-docs-all: gen-doc gen-docs-md

.PHONY: gen-models
gen-models:
	@if ! command -v sqlc >/dev/null 2>&1; then \
		echo "sqlc is not installed."; \
		echo "Please install it first by running:"; \
		echo "  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"; \
		exit 1; \
	fi
	sqlc generate

test-unit:
	@echo "Running unit tests..."
	go test -v ./...

# Test coverage target
.PHONY: test-cover
test-cover:
	@echo "Running unit tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

test-int:
	@echo "Running integration tests..."
	go test -v -tags=integration ./tests/integration/

.PHONY: build
build:
	@echo "Building project..."
	@mkdir -p $(BINARY_DIR)
	@\
	VERSION=$$(./cmd/version/version.sh); \
	BUILD_TIME=$$(date -u "+%Y-%m-%dT%H:%M:%SZ"); \
	COMMIT_HASH=$$(git rev-parse --short HEAD 2>/dev/null || echo "N/A"); \
	echo "	Version: $$VERSION"; \
	echo "	Build Time: $$BUILD_TIME"; \
	echo "	Commit Hash: $$COMMIT_HASH"; \
	go build -ldflags "-X '$(VERSION_PKG).Version=$$VERSION' -X '$(VERSION_PKG).BuildTime=$$BUILD_TIME' -X '$(VERSION_PKG).Commit=$$COMMIT_HASH'" -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_SRC)
	#go build -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_SRC)

# Minimal build target
.PHONY: build-min
min-build:
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_SRC)

# Update all packages
.PHONY: update
update:
	go get -u ./...
	go mod tidy
	go mod verify

# Clean target
.PHONY: clean
clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

# Start target
.PHONY: start
start:
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: local-db-up
local-db-up:
	@mkdir -p .devcontainer/db  # Create directory if not exists
	docker-compose -f $(DB_LOCAL_COMPOSE) up -d
	docker ps -f name=cushon-local-db

.PHONY: local-db-dn
local-db-dn:
	docker-compose -f $(DB_LOCAL_COMPOSE) down

.PHONY: local-db-destroy
local-db-destroy:
	docker-compose down -v

.PHONY: migrate-install
migrate-install:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then echo "Usage: make migrate-create name=<migration_name>"; exit 1; fi
	migrate create -ext sql -dir $(DB_MIGRATIONS_DIR) -seq $(name)

.PHONY: migrate-genesis
migrate-genesis:
	cd ./database/design && chmod +x ./genesis.sh && ./genesis.sh
	cd ../../
	make migrate-up

.PHONY: migrate-up
migrate-up:
	migrate -path $(DB_MIGRATIONS_DIR) -database "$(DB_LOCAL_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path $(DB_MIGRATIONS_DIR) -database "$(DB_LOCAL_URL)" down

# Run target (builds and runs the project after starting the database and running migrations)
.PHONY: run
run: local-db-up build migrate-up
	./$(BUILD_DIR)/$(BINARY_NAME)
