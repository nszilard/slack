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
NEW_VERSION ?= "0.2.0-pre-release"
LATEST_VERSION=$$(git describe --abbrev=0 --tags)

#--------------------------------------------------
# Targets
#--------------------------------------------------
.PHONY: bootstrap
bootstrap: ## Downloads and cleans up all dependencies
	@go mod tidy
	@go mod download

.PHONY: fmt
fmt: ## Formats go files
	@echo "==> Formatting files..."
	@gofmt -w -s $(GO_FILES)
	@echo ""

.PHONY: check
check: ## Checks code for linting/construct errors
	@echo "==> Checking if files are well formatted..."
	@gofmt -l $(GO_FILES)
	@echo ""
	@echo "==> Checking if files pass go vet..."
	@go list -f '{{.Dir}}' ./... | xargs go vet;
	@echo ""

.PHONY: test
test: check ## Runs all tests
	@echo "==> Running tests..."
	@go test -v --race $(TEST) -parallel=20
	@echo ""

.PHONY: coverage
coverage: ## Runs code coverage
	@mkdir -p .target/coverage
	@go test --p=1 $(TEST) -coverprofile=.target/coverage/cover.out -covermode=atomic

.PHONY: show-coverage
show-coverage: coverage ## Shows code coverage report in your web browser
	@go tool cover -html=.target/coverage/cover.out

.PHONY: dev
dev: fmt check ## Builds a local dev version
	@go build -ldflags "-X 'github.com/nszilard/slack/cmd.appVersion=${LATEST_VERSION}-dev'" -o .target/local/${BINARY}
	@go install -ldflags "-X 'github.com/nszilard/slack/cmd.appVersion=${LATEST_VERSION}-dev'"

.PHONY: package
package: clean bootstrap check ## Builds a production version
	@env GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/nszilard/slack/cmd.appVersion=${NEW_VERSION}'" -o .target/linux_amd64/${BINARY}

.PHONY: docs
docs: dev ## Generates markdown documentation
	@.target/local/${BINARY} docs

.PHONY: clean
clean: ## Cleans up temporary and compiled files
	@echo "==> Cleaning up ..."
	@rm -rf .target
	@echo "    [✓]"
	@echo ""

.PHONY: help
help: ## Shows available targets
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'
