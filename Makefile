PROJECT_NAME:=Padio

.PHONY: all
all: help

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build application binaries
	go build -race -o ./bin/server ./cmd/server

.PHONY: run
run: ## run server application binaries
	go run ./cmd/server

.PHONY: run-player
run-player: ## run player application binaries
	go run ./cmd/player

.PHONY: mod
mod: ## Ensured Go package dependencies
	go mod tidy

.PHONY: test
test: ## Run unit tests
	go test -v -race ./cmd/... 
