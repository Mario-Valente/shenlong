.PHONY: all build clean run

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=shenlong

all: test build

build: 
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

run: build
	./bin/$(BINARY_NAME)

test: 
	$(GOTEST) -v ./...