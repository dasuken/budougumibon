.PHONY: help build build-local up down logs ps tst
.DEFAULT_GOAL := help

migrate:
	psqldef  -U todo -p 5433 -h 127.0.0.1 -W todo todo  < ./_tools/postgresql/schema.sql

build: ## Build docker image to deploy
	docker build -t hoge25/gotodo:${DOCKER_TAG} \
		-- target deploy ./

build-local: ## Build docker image to local deployment
	docker compose build --no-cache

up:
	docker compose up -d

logs:
	docker compose logs -f

ps:
	docker compose ps

test:
	go test -race -shuffle=on ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'