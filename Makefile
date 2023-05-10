PROJECT_NAME=app
MAIN_FILE=app.go
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')

.PHONY: all dep listvar build run clean help

all: build

dep: ## Get the dependencies
	@go mod tidy

listvar: ## List var
	@echo BUILD_BRANCH=$(BUILD_BRANCH) && \
	echo BUILD_COMMIT=$(BUILD_COMMIT) && \
	echo BUILD_TIME=$(BUILD_TIME) && \
	echo BUILD_GO_VERSION=$(BUILD_GO_VERSION)

build: dep ## Build the binary file
	@go build -ldflags "-s -w" -o dist/app.exe $(MAIN_FILE)

run: ## Run Develop server
	@go run $(MAIN_FILE)

clean: ## Remove previous build
	@rm -f dist/*

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
