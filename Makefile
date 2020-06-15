# Makefile to build go-sdk-template

all: build unittest lint tidy

travis-ci: build alltest lint tidy

build:
	go build ./...

unittest:
	go test `go list ./... | grep -v samples`

alltest:
	go test `go list ./... | grep -v samples` -v -tags=integration

lint:
	golangci-lint run

tidy:
	go mod tidy
