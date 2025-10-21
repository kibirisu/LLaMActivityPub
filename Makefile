BACKEND_DIR := server
FRONTEND_DIR := web

.PHONY: dev
dev:
	@$(MAKE) -j2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend:
	cd $(BACKEND_DIR) && air

.PHONY: dev-frontend
dev-frontend:
	cd $(FRONTEND_DIR) && pnpm run dev
