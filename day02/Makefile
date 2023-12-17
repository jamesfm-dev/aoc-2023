input := ""

default: help

help: ## USAGE: make input="path/to/input.txt" <command>
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: c
c: ## Run in C, USAGE: make input="path/to/input.txt" c
	@gcc -Wall -Werror ./c/main.c -o ./c/a.out && ./c/a.out $(input)

.PHONY: go
go: ## Run in Go, USAGE: make input="path/to/input.txt" go
	@go run ./go/main.go $(input)
