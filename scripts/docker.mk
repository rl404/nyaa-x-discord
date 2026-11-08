# Compose command.
COMPOSE ?= docker compose

# Compose files helper.
compose = $(COMPOSE) -f deployment/compose.$(1).yml -p nxd-$(2)

.PHONY: docker-lint docker-build docker-bot docker-cron-check docker-stop

##@ Docker Tools

docker-lint: ## Run linter in docker.
	@$(call compose,lint,lint) --profile lint up
	@$(call compose,lint,lint) --profile lint rm -f

##@ Docker Build & Run

docker-build: ## Build docker image then prune the builder image.
	@$(call compose,build,build) --profile build build 
	@docker image prune -f --filter label=stage=nxd_builder

docker-bot: ## Start bot container.
	@$(call compose,bot,bot) up -d
	@$(call compose,bot,bot) logs --follow --tail 20

docker-cron-check: ## Start cron check container.
	@$(call compose,cron-check,cron-check) up
	@$(call compose,cron-check,cron-check) rm -f

docker-stop: ## Stop running container.
	@$(call compose,bot,bot) stop
	@$(call compose,bot,bot) rm -f
