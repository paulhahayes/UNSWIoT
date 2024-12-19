CLI_NAME := noRun
CLI_DIR := backend/src/noRun
CLI_BINARY := $(CLI_DIR)/$(CLI_NAME)

BACKEND_NAME := brun
BACKEND_DIR := backend/src/app
BACKEND_BINARY := $(BACKEND_DIR)/$(BACKEND_NAME)

DB_NAME := IoT
DB_DIR := backend/src/internal/db

INSTALL_PATH := /usr/local/bin

.PHONY: all cli backend db install-cli install-backend clean

all: cli backend

cli:
	@echo "Building CLI..."
	@echo "Waiting..."
	@cd $(CLI_DIR) && go build
	@cd $(CLI_DIR) && go install
	@echo "CLI binary built at $(CLI_DIR)/$(CLI_NAME)"

backend:
	@echo "Building backend..."
	@cd $(BACKEND_DIR) && go build -o $(BACKEND_NAME)
	@sudo mv $(BACKEND_DIR)/$(BACKEND_NAME) $(INSTALL_PATH)/$(BACKEND_NAME)
	@echo "Backend binary built at $(BACKEND_DIR)/$(BACKEND_NAME)"

db:
	@sqlite3 $(DB_DIR)/$(DB_NAME).db < $(DB_DIR)/$(DB_NAME).sql

db-test:
	@sqlite3 $(DB_DIR)/Test_IoT.db < $(DB_DIR)/$(DB_NAME).sql
clean:
	@echo "Cleaning up..."
	@sudo rm -f $(INSTALL_PATH)/$(CLI_NAME)
	@sudo rm -f $(INSTALL_PATH)/$(BACKEND_NAME)