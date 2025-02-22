# Default recipe to display help
default:
    @just --list

# Install just command itself (for Linux)
install-just:
    #!/bin/bash
    TEMP_DEB="$(mktemp)"
    curl -LSs 'https://github.com/casey/just/releases/download/1.14.0/just-1.14.0-x86_64-unknown-linux-musl.tar.gz' -o "$TEMP_DEB"
    tar -xzf "$TEMP_DEB"
    sudo mv just /usr/local/bin/
    rm "$TEMP_DEB"
    @echo "just installed successfully"

# Install all tools and dependencies
install-all: install-just install

# Install dependencies
install:
    go mod download
    go install github.com/objectbox/objectbox-go/cmd/objectbox-gogen@latest
    go install github.com/swaggo/swag/cmd/swag@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Generate ObjectBox and Swagger files
generate:
    go generate ./...
    swag init

# Run the application locally
run:
    go run main.go

# Build the application
build:
    CGO_ENABLED=1 go build -o bin/api main.go

# Run tests
test:
    go test -v ./...

# Run tests with coverage
test-coverage:
    go test -v -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report generated at coverage.html"

# Clean build artifacts
clean:
    rm -rf bin/
    rm -rf objectbox/
    rm -f coverage.out coverage.html

# Docker commands
docker-build:
    docker-compose build

docker-up:
    docker-compose up -d

docker-down:
    docker-compose down

docker-logs:
    docker-compose logs -f

docker-restart: docker-down docker-up

# Development helpers
swagger:
    swag init
    @echo "Swagger docs updated at http://localhost:8080/swagger/index.html"

lint:
    golangci-lint run

# Database commands
db-clean:
    rm -rf objectbox/

# Monitoring commands
metrics:
    @echo "Prometheus: http://localhost:9090"
    @echo "Grafana: http://localhost:3000 (admin/admin)"

# Full development setup
dev-setup: install generate

# Full deployment
deploy: generate docker-build docker-up
    @echo "API running at http://localhost:8080"
    @echo "Swagger docs at http://localhost:8080/swagger/index.html"
    @echo "Prometheus at http://localhost:9090"
    @echo "Grafana at http://localhost:3000"

# Update dependencies
update-deps:
    go get -u ./...
    go mod tidy

# Health check
health:
    curl -f http://localhost:8080/health

# Watch for file changes and restart (requires watchexec)
watch:
    watchexec -r -e go -- "go run main.go"

# Format all Go files
fmt:
    go fmt ./...

# Verify all dependencies
verify:
    go mod verify

# Show all running containers
ps:
    docker-compose ps

# Check API endpoints
check-api: health
    curl -s http://localhost:8080/api/v1/activities | jq
    curl -s http://localhost:8080/api/v1/stats | jq

# Full cleanup
cleanup: docker-down clean db-clean
    docker system prune -f

# Grid commands
grid-start:
    docker-compose up grid -d
    @echo "Grid UI available at http://localhost:8082"

grid-stop:
    docker-compose stop grid

grid-logs:
    docker-compose logs -f grid

grid-clean:
    docker-compose rm -f grid
    docker volume rm go-rest-api_grid-data 