# Makefile to build the project
GO=go
LINT=golangci-lint
GOSEC=gosec

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint scan-gosec

test:
	${GO} test ./...

test-cov:
	${GO} test ./... ${COVERAGE}

test-int:
	${GO} test ./... -tags=integration

test-int-cov:
	${GO} test ./... -tags=integration ${COVERAGE}

lint:
	${LINT} run --build-tags=integration,examples --timeout 120s

scan-gosec:
	${GOSEC} ./...

tidy:
	${GO} mod tidy
