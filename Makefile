BREW_PREFIX  ?= $(shell brew --prefix)
DATABASE_URL ?= "postgres://$(USER)@localhost/awaik_test?sslmode=disable"

test:
	@DATABASE_URL=$(DATABASE_URL) go test -race -timeout 1s ./...

test-setup: $(BREW_PREFIX)/bin/migrate
	migrate -path migrations/ -database $(DATABASE_URL) up

build:
	@go build -o ./bin/awaik ./cmd/awaik

install:
	@go install ./cmd/awaik

lint:
	@golangci-lint run
	buf lint

deps:
	brew install golangci-lint golang-migrate

generate:
	buf generate

tui:
	@go run ./cmd/awaik --tui

$(BREW_PREFIX)/bin/migrate:
	@brew install golang-migrate

.PHONY: all build deps lint test