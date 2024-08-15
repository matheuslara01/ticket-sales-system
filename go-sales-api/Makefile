.PHONY: all clean swag init run

SWAGGER_DIR = ./docs
CMD_DIR = ./cmd/events
HANDLER_DIR = ./internal/events/infra/http

BINARY_NAME = events

DB_USER = test_user
DB_PASSWORD = test_password
DB_HOST = localhost
DB_PORT = 3306
DB_NAME = test_db

all: swag build run

clean:
	rm -rf $(SWAGGER_DIR)
	rm -rf $(BINARY_NAME)

swag: 
	swag init --output docs --dir ./cmd/events,./internal/events/infra/http,./internal/events/usecase

init:
	mysql -u$(DB_USER) -p$(DB_PASSWORD) -h$(DB_HOST) -P$(DB_PORT) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"

build:
	go build -o $(BINARY_NAME) $(CMD_DIR)/main.go

run: build
	./$(BINARY_NAME)

dev: swag build run