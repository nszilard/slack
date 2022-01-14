#!make
#----------------------------------------
# Settings
#----------------------------------------
.DEFAULT_GOAL := help

#--------------------------------------------------
# Variables
#--------------------------------------------------
BINARY="slack"
TEST?=$$(go list ./...)
GO_FILES?=$$(find . -name '*.go')

#--------------------------------------------------
# Targets
#--------------------------------------------------
bootstrap: ## Downloads and cleans up all dependencies
	@go mod tidy
	@go mod download

fmt: ## Formats go files
	@echo "==> Formatting files..."
	@gofmt -w -s $(GO_FILES)
	@echo ""

check: ## Checks code for linting/construct errors
	@echo "==> Checking if files are well formatted..."
	@gofmt -l $(GO_FILES)
	@echo ""
	@echo "==> Checking if files pass go vet..."
	@go list -f '{{.Dir}}' ./... | xargs go vet;
	@echo ""

test: check ## Runs all tests
	@echo "==> Running tests..."
	@go test -v --race $(TEST) -parallel=20
	@echo ""

coverage: ## Runs code coverage
	@mkdir -p .target/coverage
	@go test --p=1 $(TEST) -coverprofile=.target/coverage/cover.out -covermode=atomic

show-coverage: coverage ## Shows code coverage report in your web browser
	@go tool cover -html=.target/coverage/cover.out

dev: fmt check ## Builds a local dev version
	@go build -o .target/local/${BINARY}
	@go install

package: bootstrap check ## Builds a production version
	@env GOOS=linux GOARCH=amd64 go build -o .target/linux_amd64/${BINARY}

docs: dev ## Generates markdown documentation
	@.target/local/${BINARY} docs

.PHONY: bootstrap fmt check test coverage show-coverage dev package docs clean help

clean: ## Cleans up temporary and compiled files
	@echo "==> Cleaning up ..."
	@rm -rf .target
	@echo "    [✓]"
	@echo ""

help: ## Shows available targets
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'
