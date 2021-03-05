# Base Go commands.
GO_CMD   := go
GO_FMT   := $(GO_CMD) fmt
GO_CLEAN := $(GO_CMD) clean
GO_BUILD := $(GO_CMD) build -mod vendor

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
	@golint `go list ./... | grep -v /vendor/`

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

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

# Build and run checker.
.PHONY: check
check: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) check

# Build and run cron.
.PHONY: cron
cron: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) cron

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD   := docker-compose
COMPOSE_PATH  := deployment/docker-compose.yml

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build: clean fmt
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=nxd_builder

# Start built docker containers.
.PHONY: docker-up
docker-up:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) logs --follow --tail 20

# Build and start docker container for the project.
.PHONY: docker
docker: docker-build docker-up

# Stop docker container.
.PHONY: docker-stop
docker-stop:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) stop