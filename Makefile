BREW_PREFIX  ?= $(shell brew --prefix)
DATABASE_URL ?= "postgres://$(USER)@localhost/uptime_test?sslmode=disable"

test:
	@DATABASE_URL=$(DATABASE_URL) go test -race -timeout 1s ./...

test-setup: $(BREW_PREFIX)/bin/migrate
	migrate -path migrations/ -database $(DATABASE_URL) up

build:
	@go build ./cmd/uptime

lint:
	@golangci-lint run

deps:
	brew install golangci-lint golang-migrate

$(BREW_PREFIX)/bin/migrate:
	@brew install golang-migrate

.PHONY: all build deps lint test