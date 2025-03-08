#@IgnoreInspection BashAddShebang
ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
CGO_ENABLED?=0

APP_NAME?=ketabdoozak

IN_DOCKER?=false

IMAGE_REPOSITORY?=ghcr.io/nasermirzaei89/ketabdoozak
IMAGE_TAG?=latest

OS_NAME=$(shell uname -s | tr '[:upper:]' '[:lower:]')
OS_ARCH=$(shell uname -m)

ifeq ($(OS_NAME),darwin)
    OS_NAME := macos
endif

SWAG_CMD=go tool swag
GOLANGCI_LINT_CMD=go tool golangci-lint
AIR_CMD=go tool air
TEMPL_CMD=go tool templ
TAGALIGN_CMD=go tool tagalign

TAILWINDCSS_VERSION=v4.0.7
TAILWINDCSS_URL="https://github.com/tailwindlabs/tailwindcss/releases/download/$(TAILWINDCSS_VERSION)/tailwindcss-$(OS_NAME)-$(OS_ARCH)"
TAILWINDCSS_CMD=$(ROOT)/bin/tailwindcss

.DEFAULT_GOAL := .default

.default: format generate-docs build lint test

.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.which-go:
	@which go > /dev/null || (echo "Install Go from https://go.dev/doc/install" & exit 1)

.PHONY: build
build: .which-go ## Build binary
ifeq ($(IN_DOCKER),false)
	make templ-generate
	make tailwindcss-build
endif
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/$(APP_NAME) $(ROOT)

.PHONY: dev
dev: ## Run dev mode
	make -j4 air-run tailwindcss-watch npm-watch templ-watch air-watch-static

.PHONY: templ-generate
templ-generate: ## Generate Templ Go files
	$(TEMPL_CMD) generate

.PHONY: templ-watch
templ-watch: ## Watch and generate Templ Go files
	$(TEMPL_CMD) generate -watch -proxy=http://localhost:8080 -proxyport 3000

air-run: .which-go
	$(AIR_CMD) \
 --build.cmd "go build -o tmp/main" --build.bin "tmp/main" --build.delay "1000" \
 --build.exclude_dir "docs,infra" \
 --build.include_ext "go,css,js" \
 --build.include_file ".env" \
 --build.stop_on_error "false" \
 --build.pre_cmd "make generate-docs"

.which-tailwindcss:
	@which $(TAILWINDCSS_CMD) > /dev/null || ( \
		echo "Downloading $(TAILWINDCSS_URL)" && \
		curl -L --create-dirs -o $(TAILWINDCSS_CMD) $(TAILWINDCSS_URL) && \
        chmod +x  $(TAILWINDCSS_CMD) \
	)

.PHONY: tailwindcss-watch
tailwindcss-watch: .which-tailwindcss  ## Watch and build style.css
	$(TAILWINDCSS_CMD) --input $(ROOT)/www/assets/style.css --output $(ROOT)/www/static/style.css --watch

.PHONY: tailwindcss-build
tailwindcss-build: .which-tailwindcss ## Build style.css
	$(TAILWINDCSS_CMD) --input $(ROOT)/www/assets/style.css --output $(ROOT)/www/static/style.css
	$(TAILWINDCSS_CMD) --input $(ROOT)/www/assets/style.css --output $(ROOT)/www/static/style.min.css --minify

.which-npm:
	@which npm > /dev/null || (echo "Install NodeJS from https://nodejs.org/en/download" & exit 1)

.PHONY: npm-watch
npm-watch: .which-npm
	npm --prefix ./www run watch

.PHONY: npm-build
npm-build: .which-npm
	npm --prefix ./www run build
	npm --prefix ./www run build:min

.PHONY: air-watch-static
air-watch-static:
	$(AIR_CMD) \
 --build.cmd "$(TEMPL_CMD) generate --notify-proxy -proxyport 3000" \
 --build.bin "true" \
 --build.delay "100" \
 --build.exclude_dir "" \
 --build.include_dir "www/static" \
 --build.include_ext "js,css"

.PHONY: format
format: .which-go ## Format files
	go mod tidy
	gofmt -s -w $(ROOT)
	$(SWAG_CMD) fmt
	$(TEMPL_CMD) fmt $(ROOT)/www/templates/
	$(TAGALIGN_CMD) -fix $(ROOT)/... || echo "tags aligned"

.PHONY: lint
lint: .which-go ## Check lint
	$(GOLANGCI_LINT_CMD) run

.PHONY: test
test: .which-go ## Run tests
	CGO_ENABLED=1 go test -race -cover $(ROOT)/...

.PHONY: generate-docs
generate-docs: .which-go ## Generate swagger files
	$(SWAG_CMD) init

.which-docker:
	@which docker > /dev/null || (echo "Install Docker from https://www.docker.com/get-started/" & exit 1)

.PHONY: docker-build
docker-build: .which-docker ## Build docker image
	docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) $(ROOT)

.PHONY: docker-push
docker-push: .which-docker ## Push docker image
	docker push $(IMAGE_REPOSITORY):$(IMAGE_TAG)
