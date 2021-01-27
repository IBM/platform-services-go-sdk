# Makefile to build the project

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: test lint tidy
travis-ci: test-cov lint tidy

test:
	go test `go list ./...`

test-cov: 
	go test `go list ./...` ${COVERAGE}

test-int:
	go test `go list ./...` -tags=integration

test-int-cov:
	go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	golangci-lint run

tidy:
	go mod tidy
