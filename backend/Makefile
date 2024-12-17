# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=bookpavilion
BINARY_UNIX=$(BINARY_NAME)_unix

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build test clean run deps setup-test test-verbose

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: setup-test
	$(GOTEST) ./...

test-verbose: setup-test
	$(GOTEST) -v ./...

clean: 
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf test_uploads

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

deps:
	$(GOMOD) download

setup-test:
	@echo "Setting up test environment..."
	@if [ ! -d "test_uploads" ]; then \
		mkdir -p test_uploads; \
	fi
	@if [ ! -f ".env.test" ]; then \
		echo "DB_HOST=localhost" > .env.test; \
		echo "DB_PORT=3306" >> .env.test; \
		echo "DB_USER=root" >> .env.test; \
		echo "DB_PASSWORD=" >> .env.test; \
		echo "DB_NAME=bookpavilion_test" >> .env.test; \
		echo "UPLOAD_DIR=./test_uploads" >> .env.test; \
	fi
	@mysql -u root -e "CREATE DATABASE IF NOT EXISTS bookpavilion_test;"
	@mysql -u root bookpavilion_test < scripts/setup_test_db.sql

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t $(BINARY_NAME) .

# Database operations
db-create:
	mysql -u root -e "CREATE DATABASE IF NOT EXISTS bookpavilion;"
	mysql -u root -e "CREATE DATABASE IF NOT EXISTS bookpavilion_test;"

# Help command
help:
	@echo "Make commands:"
	@echo "make build         - Build the application"
	@echo "make test         - Run tests"
	@echo "make test-verbose - Run tests with verbose output"
	@echo "make clean        - Clean build files"
	@echo "make run          - Build and run the application"
	@echo "make deps         - Download dependencies"
	@echo "make setup-test   - Set up test environment"
	@echo "make db-create    - Create development and test databases"
