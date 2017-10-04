BINARY = Hacksite

ENVIRONMENT_FILE=environments/dev.env.json

BUILD_FLAGS=-ldflags "-X main.envFile=${ENVIRONMENT_FILE}"

.DEFAULT_GOAL: all

.PHONY: all
all: setup server web

.PHONY: local
local: setup buildLocal web 

.PHONY: setup
setup: setup-go setup-web

.PHONY: setup-go
setup-go:
	@echo "> Getting go dependencies"
	@./scripts/setup-go.sh

.PHONY: setup-web
setup-web:
	@echo "> Getting web dependencies"
	@./scripts/setup-web.sh

.PHONY: buildLocal
buildLocal: server
	@echo "> Generating local certificates"
	./scripts/generateCerts.sh

.PHONY: server
server:
	@echo "> Building Server"
	@go build ${PROD_BUILD_FLAGS} -o ${BINARY} server/cmd/server.go

.PHONY: web
web:
	@./scripts/buildWeb.sh

.PHONY: help
help:
	@echo "default	- Builds server and web client for production"
	@echo "local	- Builds server for local use and web client"
