DOCKER_IMAGE := gitdump:latest
GO_BINARY_NAME := gitdump
RELEASE_DIR := bin
OSX_BINARY_DIR := $(RELEASE_DIR)/osx
LINUX_BINARY_DIR := $(RELEASE_DIR)/linux
WINDOWS_BINARY_DIR := $(RELEASE_DIR)/windows

.PHONY: build dev binary releases clean

build:
	docker build -t $(DOCKER_IMAGE) .

dev: build
	docker run -it --rm -v $(CURDIR):/usr/src/app/go/src/github.com/mrauer/gitdump $(DOCKER_IMAGE)

binary:
	env GOOS=linux GOARCH=amd64 go build -o $(GO_BINARY_NAME)

releases: clean binary
	env GOOS=darwin GOARCH=amd64 go build -o $(OSX_BINARY_DIR)/$(GO_BINARY_NAME)
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX_BINARY_DIR)/$(GO_BINARY_NAME)
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BINARY_DIR)/$(GO_BINARY_NAME).exe

clean:
	rm -rf $(RELEASE_DIR)
