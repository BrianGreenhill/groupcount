all: build test fmt vet

.PHONY: all

build:
	go build ./...

test:
	go test ./...

fmt:
	@go fmt ./...

vet:
	go vet ./...
