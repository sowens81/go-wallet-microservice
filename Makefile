
# Makefile - common development commands
.PHONY: build run test lint clean

# Build the application
build:
    go build -o bin/walletapi ./cmd/walletapi

# Run the application
run:
    go run ./cmd/walletapi

# Run all tests with coverage
test:
    go test -v -race -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

# Run linters
lint:
    golangci-lint run

# Clean build artifacts
clean:
    rm -rf bin/
    rm -f coverage.out coverage.html

# Download dependencies
deps:
    go mod download
    go mod tidy

# Build for multiple platforms
build-all:
    GOOS=linux GOARCH=amd64 go build -o bin/walletapi-linux-amd64 ./cmd/walletapi
    GOOS=darwin GOARCH=arm64 go build -o bin/walletapi-darwin-arm64 ./cmd/walletapi
    GOOS=windows GOARCH=amd64 go build -o bin/walletapi-windows-amd64.exe ./cmd/walletapi