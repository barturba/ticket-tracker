include .env
# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/ticket-tracker: run the cmd/ticket-tracker application
.PHONY: run/ticket-tracker
run/ticket-tracker:
	go run ./cmd/ticket-tracker -db-dsn=${DATABASE_URL}

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/ticket-tracker: build the cmd/ticket-tracker application
.PHONY: build/ticket-tracker
build/ticket-tracker:
	@echo 'Building cmd/ticket-tracker ...'
	go build -ldflags="-s -w" -o=./bin/ticket-tracker ./cmd/ticket-tracker
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o=./bin/darwin_arm64/ticket-tracker ./cmd/ticket-tracker