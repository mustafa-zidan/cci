# Variables
PACKAGES := $(shell go list ./...)
SOURCES  := $(shell ls *.go | grep -v _test.go)
RACE     := $(shell test $$(go env GOARCH) != "amd64" || echo "-race")
VERSION  := $(shell cat VERSION)

.PHONY: help run test benchmark build vet

help: ## Show this help.
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application.
	@echo "Compiling version $(VERSION)..."
	@go run -o cci $(SOURCES) -i=$(i) -o=$(o)

test: ## Run tests (excluding integration tests).
	@go test $(RACE) $(PACKAGES)

benchmark: ## Run benchmark tests.
	@go test --bench=^Benchmark.*$$ -benchmem -benchtime=10000000x $(PACKAGES)

build: ## Build the application.
	@echo "Compiling version $(VERSION)..."
	@mkdir -p bin
	@go build -o bin/cci $(SOURCES) -i=$(i) -o=$(o)
	@echo "All done! The binaries are in ./bin. Check it out!"

vet: ## Run go vet.
	@echo "Running go vet..."
	@test -z "$$(go vet $(PACKAGES) 2>&1 | grep -v 'composite literal uses unkeyed fields' | tee /dev/stderr)" || (echo "go vet found issues!" && exit 1)