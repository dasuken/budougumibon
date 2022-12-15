.PHONY: help build build-local up down logs ps tst
.DEFAULT_GOAL := help

postgres:
	docker run --name postgres -p 5432:5432  -e POSTGRES_USER=todo -e POSTGRES_PASSWORD=todo -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=todo --owner=todo todo

dropdb:
	docker exec -it postgres dropdb simple_bank

execpsql:
	docker exec -it postgres /usr/local/bin/psql -U todo -c "$(C)"

migrate: ## db container port 55432. if you use psql well known port, change this value.
	psqldef  -U todo -p 55432 -h 127.0.0.1 -W todo todo  < ./_tools/postgresql/schema.sql

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