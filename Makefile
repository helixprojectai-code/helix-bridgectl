BINARY_NAME = helix-bridgectl
VERSION ?= 0.1.0

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/helix-bridgectl

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: install
install: build
	sudo mv $(BINARY_NAME) /usr/local/bin/

.PHONY: run
run: build
	./$(BINARY_NAME)
