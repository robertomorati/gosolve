.PHONY: test build run docker-build docker-run down

test:
	@echo "Running tests..."
	go test -cover -v ./internal/...

coverage:
	@echo "Generating test coverage report..."
	go test -coverprofile=coverage.out -v ./internal/...
	go tool cover -func=coverage.out

up:
	@echo "Building and starting the application with Docker..."
	docker-compose up --build -d

down:
	@echo "Stopping and removing Docker containers..."
	docker-compose down
