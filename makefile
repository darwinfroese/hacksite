BINARY = Hacksite

.DEFAULT_GOAL: all

.PHONY: all
all: web server

.PHONY: server
server:
	@echo "> Building Server"
	go build -o ${BINARY} server/cmd/server.go

.PHONY: web
web:
	./buildWeb.sh

.PHONY: help
help:
	@echo "default	- Builds server and web client"
	@echo "server	- Builds the server"
	@echo "web 	- Builds the web client"