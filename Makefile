APP_NAME := borg
FRONTEND_DIR := web
DIST_DIR := $(FRONTEND_DIR)/dist
BIN_DIR := $(PWD)/bin
TOOLS := air sqlc
GO_BUILD_CMD := go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

NODE_MODULES := $(FRONTEND_DIR)/node_modules
LOCKFILE := $(FRONTEND_DIR)/pnpm-lock.yaml
PACKAGE_JSON := $(FRONTEND_DIR)/package.json

export PATH := $(BIN_DIR):$(PATH)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: tools
tools: $(BIN_DIR)/sqlc $(BIN_DIR)/air

$(BIN_DIR)/sqlc: | $(BIN_DIR)
	@echo Installing sqlc...
	@GOBIN=$(BIN_DIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

$(BIN_DIR)/air: | $(BIN_DIR)
	@echo Installing air...
	@GOBIN=$(BIN_DIR) go install github.com/air-verse/air@latest

$(NODE_MODULES): $(LOCKFILE) $(PACKAGE_JSON)
	cd $(FRONTEND_DIR) && pnpm install --frozen-lockfile

.PHONY: run
run: build-backend 
	$(BIN_DIR)/$(APP_NAME)

.PHONY: build
build: build-backend build-frontend

.PHONY: build-backend
build-backend:
	$(GO_BUILD_CMD)

.PHONY: build-frontend
build-frontend: $(NODE_MODULES)
	cd $(FRONTEND_DIR) && pnpm build

.PHONY: dev
dev:
	@$(MAKE) -j2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend: tools
	@echo Starting dev server...
	@air -build.cmd "$(GO_BUILD_CMD)" -build.bin "$(BIN_DIR)/$(APP_NAME)" -build.exclude_dir "bin,web" -build.post_cmd "rmdir tmp"

.PHONY: dev-frontend
dev-frontend: $(NODE_MODULES)
	@echo Starting dev react app...
	@cd $(FRONTEND_DIR) && pnpm dev

.PHONY: run-compose
run-compose:
	docker compose up -d

.PHONY: run-dev-db
run-dev-db:
	docker compose up -d db

.PHONY: gen-sql
gen-sql: tools
	@echo Generating sqlc modules...
	@sqlc generate -f .sqlc.yaml

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
	rm -rf $(NODE_MODULES)
	rm -rf $(BIN_DIR)
