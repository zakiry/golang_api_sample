NAME     := sample_app
VERSION  := $(shell git describe --tags)
REVISION := $(shell git rev-parse --short HEAD)
SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME) cmd/app.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: deps
deps:
	dep ensure -v

.PHONY: proto
proto:
	protoc -I=proto/ --twirp_out=./proto --go_out=./proto proto/*.proto
