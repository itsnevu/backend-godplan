.PHONY: swagger run build test lint clean help

help:
	@echo "Available targets:"
	@echo "  make swagger  - Generate Swagger documentation"
	@echo "  make run      - Run the application"
	@echo "  make build    - Build the binary"
	@echo "  make test     - Run tests with coverage"
	@echo "  make lint     - Run golangci-lint"
	@echo "  make clean    - Clean build artifacts"

swagger:
	swag init -g cmd/main.go -o docs

run:
	go run cmd/main.go

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/api cmd/main.go

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
	rm -f *.exe *.exe~
	rm -f coverage.out coverage.html
	go clean
	go clean -testcache
