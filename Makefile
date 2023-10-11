# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Build parameters
BINARY_NAME=election-api
BUILD_PATH=target/builds

# List of operating systems
OS_LIST=darwin linux windows

all: build

build:
	$(foreach os, $(OS_LIST), \
		$(GOTEST) -v ./... && $(GOBUILD) -o $(BUILD_PATH)/$(os)/$(BINARY_NAME) -v;)

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_PATH)

run:
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -v ./...
	$(BUILD_PATH)/$(BINARY_NAME)

.PHONY: all build clean run
