#
# Makefile for developing using Docker
#

# Overridable env vars
DOCKER_MOUNTS ?= -v "$(CURDIR)":/go/src/PockerView
DOCKER_SHELL ?= ash
DOCKER_CONTAINER_NAME ?=
DOCKER_GO_BUILD_CACHE ?= y

DEV_DOCKER_IMAGE_NAME = docker-dev$(IMAGE_TAG)
BINARY_NATIVE_IMAGE_NAME = docker-native$(IMAGE_TAG)
CROSS_IMAGE_NAME = docker-cross$(IMAGE_TAG)
CACHE_VOLUME_NAME := docker-dev-cache
ifeq ($(DOCKER_GO_BUILD_CACHE),y)
DOCKER_MOUNTS += -v "$(CACHE_VOLUME_NAME):/root/.cache/go-build"
endif
VERSION = $(shell cat VERSION)
ENVVARS = -e VERSION=$(VERSION)

# build docker image
.PHONY: build_docker_image
build_docker_image:
	cat ./dockerfiles/Dockerfile.dev | docker build ${DOCKER_BUILD_ARGS} -t $(DEV_DOCKER_IMAGE_NAME) -

.PHONY: build_cross_image
build_cross_image:
	cat ./dockerfiles/Dockerfile.cross | docker build ${DOCKER_BUILD_ARGS} -t $(CROSS_IMAGE_NAME) -

.PHONY: build_binary_native_image
build_binary_native_image:
	cat ./dockerfiles/Dockerfile.binary | docker build -t $(BINARY_NATIVE_IMAGE_NAME) -

DOCKER_RUN_NAME_OPTION := $(if $(DOCKER_CONTAINER_NAME),--name $(DOCKER_CONTAINER_NAME),)
DOCKER_RUN := docker run --rm $(ENVVARS) $(DOCKER_MOUNTS) $(DOCKER_RUN_NAME_OPTION)

binary: build_binary_native_image Gopkg.toml ## build the source
	$(DOCKER_RUN) $(BINARY_NATIVE_IMAGE_NAME)

build: binary ## alias for binary

.PHONY: clean
clean: build_docker_image ## clean build artifacts
	$(DOCKER_RUN) $(DEV_DOCKER_IMAGE_NAME) make clean
	docker volume rm -f $(CACHE_VOLUME_NAME)

.PHONY: binary-windows
binary-windows: build_cross_image Gopkg.toml ## build the source for Windows
	$(DOCKER_RUN) $(CROSS_IMAGE_NAME) make $@

.PHONY: binary-osx
binary-osx: build_cross_image Gopkg.toml ## build the source for macOS
	$(DOCKER_RUN) $(CROSS_IMAGE_NAME) make $@

.PHONY: dev
dev: build_docker_image ## start a build container for in-container development
	$(DOCKER_RUN) -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		$(DEV_DOCKER_IMAGE_NAME) $(DOCKER_SHELL)

.PHONY: set-vendor
set-vendor: build_docker_image ## set vendor and Gopkg files
	$(DOCKER_RUN) -it $(DEV_DOCKER_IMAGE_NAME) dep init

.PHONY: update-vendor
update-vendor: build_docker_image Gopkg.toml ## update vendor and Gopkg files
	$(DOCKER_RUN) -it $(DEV_DOCKER_IMAGE_NAME) dep ensure

Gopkg.toml: set-vendor


.PHONY: fmt
fmt: ## run gofmt
	$(DOCKER_RUN) $(DEV_DOCKER_IMAGE_NAME) make fmt

.PHONY: help
help: ## print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: debug
debug: ## print this help
	@echo $(IMAGE_TAG)
