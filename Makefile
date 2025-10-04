.PHONY: test build clean lint fmt vet coverage help

help:
	@echo "Available targets:"
	@echo "  test        - Run all tests"
	@echo "  build       - Build the project"
	@echo "  clean       - Remove build artifacts"
	@echo "  lint        - Run golangci-lint"
	@echo "  fmt         - Format code with gofmt"
	@echo "  vet         - Run go vet"
	@echo "  coverage    - Run tests with coverage report"
	@echo "  install     - Install dependencies"
	@echo "  help        - Show this help message"

test:
	go test -v ./tests/...

build:
	go build -v ./...

clean:
	go clean
	rm -f coverage.out coverage.html

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

vet:
	go vet ./...

coverage:
	go test -coverprofile=coverage.out -coverpkg=./... ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

install:
	go mod download
	go mod tidy

all: fmt vet test build
