.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	CGO_ENABLED=1 go test -v -race -timeout 10s ./...

.DEFAULT_GOAL := build