all:
.PHONY: all

test:
	go test -cover ./...
.PHONY: test

lint:
	golangci-lint run --fast
.PHONY: lint
