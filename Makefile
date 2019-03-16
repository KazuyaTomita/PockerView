
_:=$(shell ./scripts/warn-outside-container $(MAKECMDGOALS))


.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: binary
binary: ## build executable for linux
	./scripts/binary

.PHONY: clean
clean: ## remove build artifacts
	rm -rf ./build/*


.PHONY: fmt
fmt: ## formatting
	go list -f {{.Dir}} ./... | xargs gofmt -w -s -d


.PHONY: debug
debug: ## for debug
	@echo "hoge"
