APP_NAME := llamap
BACKEND_DIR := server
FRONTEND_DIR := web
BIN_DIR := $(PWD)/bin
TOOLS := air sqlc

NODE_MODULES := $(FRONTEND_DIR)/node_modules
LOCKFILE := $(FRONTEND_DIR)/pnpm-lock.yaml
PACKAGE_JSON := $(FRONTEND_DIR)/package.json

export PATH := $(BIN_DIR):$(PATH)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: tools
tools: $(BIN_DIR)/sqlc $(BIN_DIR)/air

$(BIN_DIR)/sqlc: | $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

$(BIN_DIR)/air: | $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install github.com/air-verse/air@latest

$(NODE_MODULES): $(LOCKFILE) $(PACKAGE_JSON)
	cd $(FRONTEND_DIR) && pnpm install --frozen-lockfile

.PHONY: dev
dev: tools
	@$(MAKE) -j2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend:
	cd $(BACKEND_DIR) && air --build.cmd "go build -o $(BIN_DIR)/$(APP_NAME) main.go" --build.bin "$(BIN_DIR)/$(APP_NAME)"

.PHONY: dev-frontend
dev-frontend: $(NODE_MODULES)
	cd $(FRONTEND_DIR) && pnpm dev

.PHONY: gen-sql
gen-sql: tools
	cd $(BACKEND_DIR)/db && sqlc generate
