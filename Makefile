# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/ticket-tracker: run the cmd/ticket-tracker application
.PHONY: run/ticket-tracker
run/ticket-tracker:
	cp .env ./cmd/ticket-tracker/.env
	go run ./cmd/ticket-tracker

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/ticket-tracker: build the cmd/ticket-tracker application
.PHONY: build/ticket-tracker
build/ticket-tracker:
	@echo 'Building cmd/ticket-tracker ...'
	go build -ldflags="-s -w" -o=./bin/ticket-tracker ./cmd/ticket-tracker
    GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o=./bin/darwin_arm64/ticket-tracker ./cmd/ticket-tracker

## build/prod/ticket-tracker/: build the cmd/ticket-tracker application for production
.PHONY: build/prod/ticket-tracker
build/prod/ticket-tracker:
	@echo 'Building cmd/ticket-tracker for production ...'
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/ticket-tracker ./cmd/ticket-tracker