.RECIPEPREFIX +=
.DEFAULT_GOAL := help
.PHONY: *

help:
	@printf "\033[33mUsage:\033[0m\n  make [target] [arg=\"val\"...]\n\n\033[33mTargets:\033[0m\n"
	@grep -E '^[-a-zA-Z0-9_\.\/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

build: ## build the containers
	@docker-compose build

up: ## Restarts and provisions the containers in the background
	@docker-compose up -d

down: ## Destroy the containers
	@docker-compose down

build_and_push: docker_build_image docker_push_image

docker_build_image:
	@docker build -t fokosun/rest-api:v1 .

docker_push_image:
	@docker push fokosun/rest-api:v1

docker_prune: prune_images prune_volumes prune_containers

prune_images: ## Remove dangling images and free up space
	@docker image prune

prune_containers: ## Remove the containers
	@docker container prune

prune_volumes: ## Removes dangling volumes
	@docker volume prune

tests: ## Run the entire test suites
	./run_tests.sh