# Makefile to build the project
GO=go
LINT=golangci-lint
GOSEC=gosec

COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint scan-gosec

test:
	${GO} test `${GO} list ./...`

test-cov:
	${GO} test `${GO} list ./...` ${COVERAGE}

test-int:
	${GO} test `${GO} list ./...` -tags=integration

test-int-cov:
	${GO} test `${GO} list ./...` -tags=integration ${COVERAGE}

lint:
	${LINT} run --build-tags=integration,examples

scan-gosec:
	${GOSEC} ./...

tidy:
	${GO} mod tidy
