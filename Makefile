.PHONY: build run test clean

# Build the application
build:
	go build -o bin/backtestd cmd/backtestd/main.go

# Run the application
run:
	go run cmd/backtestd/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download
