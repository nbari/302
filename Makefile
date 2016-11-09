.PHONY: all get test clean build cover

GO ?= go
BIN_NAME=302
VERSION=$(shell git describe --tags --always)

all: clean build

get:
	${GO} get

build: get
	${GO} build -ldflags "-X main.version=${VERSION}" -o ${BIN_NAME} cmd/302/main.go;

clean:
	@rm -rf ${BIN_NAME} *.out build 302.db

test: get
	${GO} test -v

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out
