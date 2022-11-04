default: help

help: ## Output available commands
	@echo "Available commands:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'


up: ## Start the web server
	go run cmd/web/main.go

