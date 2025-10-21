
.PHONY: dev
dev: build
	./helix-bridgectl

.PHONY: watch
watch:
	@while true; do \
		make build; \
		inotifywait -q -e modify -r . --exclude '\.git/'; \
	done
