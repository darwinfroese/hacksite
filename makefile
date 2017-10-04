BINARY = Hacksite

LOCAL_ENVIRONMENT_FILE=environments/dev.env.json
PROD_ENVIRONMENT_FILE=environments/prod.env.json

LOCAL_BUILD_FLAGS=-ldflags "-X main.envFile=${LOCAL_ENVIRONMENT_FILE}"
PROD_BUILD_FLAGS=-ldflags "-X main.envFile=${PROD_ENVIRONMENT_FILE}"

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
buildLocal: buildLocalServer
	@echo "> Generating local certificates"
	@./scripts/generateCerts.sh

.PHONY: buildLocalServer
buildLocalServer:
	@echo "> Building local server"
	@echo "> Injecting prod file: ${LOCAL_ENVIRONMENT_FILE}"
	@go build ${LOCAL_BUILD_FLAGS} -o ${BINARY} server/cmd/server.go 

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
