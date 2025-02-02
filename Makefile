PACKAGES := $(shell go list ./...)
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))
VERSION := $(shell cat VERSION)

help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run:
	@echo "Compiling version ${VERSION}... "
	@go run -o cci `ls *.go | grep -v _test.go` -i=$(i) -o=$(o)

test:           ## Run tests, except integration tests
	@go test ${RACE} ${PACKAGES}

benchmark:      ## Run benchmark tests
	@go test --bench=^Benchmark.*$$ -benchmem -benchtime 10000000x ${PACKAGES}

build:          ## Build package for multiple OSs
	@echo "Compiling version ${VERSION}... "
	@mkdir -p ./bin
	@go build -o cci `ls *.go | grep -v _test.go` -i=$(i) -o=$(o)
	@echo "All done! The binaries are in ./bin Check it out!"

vet:            ## Run go vet
	@test -z "$$(go vet ${PACKAGES} 2>&1 | grep -v '*composite literal uses unkeyed fields|exit status 0)' | tee /dev/stderr)"
