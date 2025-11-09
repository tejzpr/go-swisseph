.PHONY: all build test clean examples install help

# Variables
BINARY_NAME=go-swisseph
GO=go
GOFLAGS=-v
EXAMPLES_DIR=examples

# Default target
all: build

# Build the library
build:
	@echo "Building Swiss Ephemeris Go library..."
	$(GO) build $(GOFLAGS)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test $(GOFLAGS)

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -cover -coverprofile=coverage.out ./swisseph
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GO) test -bench=. -benchmem ./swisseph

# Build all examples
examples: build
	@echo "Building examples..."
	@mkdir -p bin
	@for dir in $(EXAMPLES_DIR)/*; do \
		if [ -d "$$dir" ]; then \
			echo "Building $$dir..."; \
			$(GO) build $(GOFLAGS) -o bin/$$(basename $$dir) $$dir; \
		fi \
	done
	@echo "Examples built in bin/ directory"

# Run basic example
run-basic: examples
	@echo "Running basic example..."
	./bin/basic

# Run natal chart example
run-natal: examples
	@echo "Running natal chart example..."
	./bin/natal_chart

# Run eclipse example
run-eclipse: examples
	@echo "Running eclipse example..."
	./bin/eclipses

# Run rise/set example
run-riseset: examples
	@echo "Running rise/set example..."
	./bin/riseset

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GO) clean
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install the library
install:
	@echo "Installing..."
	$(GO) install ./swisseph

# Show help
help:
	@echo "Available targets:"
	@echo "  make build         - Build the library"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make bench         - Run benchmarks"
	@echo "  make examples      - Build all examples"
	@echo "  make run-basic     - Run basic example"
	@echo "  make run-natal     - Run natal chart example"
	@echo "  make run-eclipse   - Run eclipse example"
	@echo "  make run-riseset   - Run rise/set example"
	@echo "  make deps          - Install dependencies"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make install       - Install the library"
	@echo "  make help          - Show this help message"

