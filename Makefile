BINARY_NAME = helix-bridgectl
VERSION ?= v0.1.0
LDFLAGS = -X github.com/helixprojectai-code/helix-bridgectl/internal/cli.version=$(VERSION) \
          -X github.com/helixprojectai-code/helix-bridgectl/internal/cli.commit=$(shell git rev-parse --short HEAD 2>/dev/null || echo dev) \
          -X github.com/helixprojectai-code/helix-bridgectl/internal/cli.date=$(shell date -u +%FT%TZ)

.PHONY: build test tidy
build:
	go build -ldflags '$(LDFLAGS)' -o $(BINARY_NAME) ./cmd/helix-bridgectl
test:
	go test ./... -v
tidy:
	go mod tidy
