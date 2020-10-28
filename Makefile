install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test -v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test

build:
	@go build -v -ldflags "-s -w -X main.version=0.0.0" -a -o bin/cline ./examples
.PHONY: build

coverage:
	# @bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage
