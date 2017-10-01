.PHONY: all deps test build

PACKAGES = $(shell go list ./... | grep -v /examples/)

all: deps test build

deps:
	go get -u github.com/spf13/viper
	go get -u github.com/op/go-logging
	go get -u github.com/devopsfaith/krakend/router/gin
	go get -u github.com/devopsfaith/krakend/router/mux

test:
	go fmt ./...
	go test -cover $(PACKAGES)
	go vet ./...

build:
	cd example/ && go build && cd .. && cp example/./example ./auth
