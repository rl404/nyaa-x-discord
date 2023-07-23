# Base Go commands.
GO_CMD   := go
GO_FMT   := $(GO_CMD) fmt
GO_CLEAN := $(GO_CMD) clean
GO_BUILD := $(GO_CMD) build

# Base golangci-lint commands.
GCL_CMD := golangci-lint
GCL_RUN := $(GCL_CMD) run

# Project executable file, and its binary.
CMD_PATH    := ./cmd/nxd
BINARY_NAME := nxd

# Default makefile target.
.DEFAULT_GOAL := bot

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@$(GCL_RUN) -D errcheck --timeout 5m

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

# Install library.
.PHONY: install
install:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.46.2
	@$(GCL_CMD) version

# Build the project executable binary.
.PHONY: build
build: clean fmt
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v .

# Build and run the discord bot.
.PHONY: bot
bot: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) bot

# Build and run cron check.
.PHONY: cron-check
cron-check: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) cron check

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD        := docker-compose
COMPOSE_BUILD      := deployment/build.yml
COMPOSE_BOT        := deployment/bot.yml
COMPOSE_CRON_CHECK := deployment/cron-check.yml

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build: clean fmt
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=nxd_builder

# Start built docker containers for bot.
.PHONY: docker-bot
docker-bot:
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) -p nxs-bot up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) -p nxs-bot logs --follow --tail 20

# Start built docker containers for cron check.
.PHONY: docker-cron-check
docker-cron-check:
	@$(COMPOSE_CMD) -f $(COMPOSE_CRON_CHECK) -p nxs-cron-check up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_CRON_CHECK) -p nxs-cron-check logs --follow --tail 20

# Stop docker container.
.PHONY: docker-stop
docker-stop:
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) -p nxs-bot stop
	@$(COMPOSE_CMD) -f $(COMPOSE_CRON_CHECK) -p nxs-cron-check stop