.DEFAULT_GOAL := help

.PHONY: gogen
gogen: ## Generate extra files
	@go generate ./...

.PHONY: godeps
godeps: ## Optimize dependencies
	@go mod tidy

.PHONY: gofmt
gofmt: ## Check code style
	@gofmt -l .

.PHONY: govet
govet: ## Run official `vet` linter
	@go vet ./...

.PHONY: golint
golint: govet ## Run all official linters
	@golint ./...

.PHONY: gocheck
gocheck: ## Run `staticcheck` linter
	@[ -x "$(shell which staticcheck)" ] || go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck ./...

.PHONY: gotest
gotest: godeps gofmt golint gocheck ## Run all tests
	@go test -vet off ./...

.PHONY: clean
clean: ## Remove build artifacts
	@go clean -cache -modcache -testcache -fuzzcache

.PHONY: help
help: # http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@echo "Usage:\n\tmake <command>\n\nThe commands are:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m\t%-16s\033[0m%s\n", $$1, $$2}'
