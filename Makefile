.PHONY: help build test lint fmt docker-up docker-down docker-logs migrate-up migrate-down

# Default target
help:
	@echo "Harasyn - Zero Trust Access Platform"
	@echo ""
	@echo "Available commands:"
	@echo "  make build         - Build all services"
	@echo "  make test          - Run all tests"
	@echo "  make lint          - Run linters"
	@echo "  make fmt           - Format code"
	@echo "  make docker-up     - Start all services with Docker Compose"
	@echo "  make docker-down   - Stop all services"
	@echo "  make docker-logs   - View Docker Compose logs"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback database migrations"
	@echo "  make dev           - Start development environment"

# Build all services
build:
	cd backend && go build -o bin/api ./cmd/api
	cd backend && go build -o bin/gateway ./cmd/gateway
	cd backend && go build -o bin/policy-worker ./cmd/policy-worker
	cd backend && go build -o bin/session-worker ./cmd/session-worker
	cd backend && go build -o bin/audit-worker ./cmd/audit-worker
	cd frontend && npm run build

# Run tests
test:
	cd backend && go test ./...
	cd frontend && npm run test

# Run linters
lint:
	cd backend && go vet ./...
	cd backend && golangci-lint run ./... || true
	cd frontend && npm run lint

# Format code
fmt:
	cd backend && go fmt ./...
	cd frontend && npm run format

# Docker Compose
docker-up:
	docker compose -f deploy/docker-compose.yml up -d --build

docker-down:
	docker compose -f deploy/docker-compose.yml down

docker-logs:
	docker compose -f deploy/docker-compose.yml logs -f

docker-ps:
	docker compose -f deploy/docker-compose.yml ps

# Database migrations (to be implemented)
migrate-up:
	@echo "Migrations not yet implemented"

migrate-down:
	@echo "Migrations not yet implemented"

# Development environment
dev: docker-up
	@echo "Services started. API: http://localhost:8080, Frontend: http://localhost:3000"

# Clean build artifacts
clean:
	rm -rf backend/bin
	rm -rf frontend/.next
	rm -rf frontend/node_modules
