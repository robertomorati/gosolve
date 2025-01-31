.PHONY: build up down restart logs backend frontend

up:
	@echo "🚀 Starting backend and frontend..."
	docker-compose up --build -d

down:
	@echo "🛑 Stopping and removing all containers..."
	docker-compose down

restart: down up

logs:
	@echo "📜 Showing logs for all services..."
	docker-compose logs -f

test:
	@echo "🧪 Running backend tests..."
	cd backend && go test -cover -v ./internal/...

coverage:
	@echo "📊 Running backend tests and generating coverage report..."
	cd backend && go test -coverprofile=coverage.out ./internal/...
	cd backend && go tool cover -func=coverage.out