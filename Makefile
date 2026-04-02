.PHONY: build install clean test vet fmt lint

BINARY=af
BUILD_DIR=./cmd/af
GO=go

build:
	$(GO) build -o $(BINARY) $(BUILD_DIR)

install:
	$(GO) install $(BUILD_DIR)

clean:
	rm -f $(BINARY)

test:
	$(GO) test ./...

vet:
	$(GO) vet ./...

fmt:
	$(GO) fmt ./...

all: fmt vet test build
