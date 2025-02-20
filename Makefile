#@IgnoreInspection BashAddShebang
ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
CGO_ENABLED?=0

APP_NAME?=ketabdoozak

IMAGE_REPOSITORY?=ghcr.io/nasermirzaei89/ketabdoozak
IMAGE_TAG?=latest

GOLANGCI_LINT_CMD=go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5
AIR_CMD=go run github.com/air-verse/air@v1.61.7
TAGALIGN_CMD=go run github.com/4meepo/tagalign/cmd/tagalign@v1.4.2

.DEFAULT_GOAL := .default

.default: format build lint test

.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.which-go:
	@which go > /dev/null || (echo "Install Go from https://go.dev/doc/install" & exit 1)

.PHONY: build
build: .which-go ## Build binary
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/$(APP_NAME) $(ROOT)

.PHONY: run
run: .which-go ## Run dev
	go run $(ROOT)

.PHONY: dev
dev: air-run ## Run dev mode

air-run: .which-go
	$(AIR_CMD) \
 --build.cmd "go build -o tmp/main" --build.bin "tmp/main" --build.delay "1000" \
 --build.include_ext "go" \
 --build.include_file ".env" \
 --build.stop_on_error "false"

.PHONY: format
format: .which-go ## Format files
	go mod tidy
	gofmt -s -w $(ROOT)
	$(TAGALIGN_CMD) -fix $(ROOT)/... || echo "tags aligned"

.PHONY: lint
lint: .which-go ## Check lint
	$(GOLANGCI_LINT_CMD) run

.PHONY: test
test: .which-go ## Run tests
	CGO_ENABLED=1 go test -race -cover $(ROOT)/...

.which-docker:
	@which docker > /dev/null || (echo "Install Docker from https://www.docker.com/get-started/" & exit 1)

.PHONY: docker-build
docker-build: .which-docker ## Build docker image
	docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) $(ROOT)

.PHONY: docker-push
docker-push: .which-docker ## Push docker image
	docker push $(IMAGE_REPOSITORY):$(IMAGE_TAG)
