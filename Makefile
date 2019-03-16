
_:=$(shell ./scripts/warn-outside-container $(MAKECMDGOALS))


.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: binary
binary: ## build executable for linux
    @echo `@pwd`
	./scripts/binary

.PHONY: fmt
fmt:
	go list -f {{.Dir}} ./... | xargs gofmt -w -s -d

vendor: vendor.conf ## check that vendor matches vendor.conf
	# TODO we install libraries

.PHONY: debug
debug: ## for debug
	@echo "hoge"
