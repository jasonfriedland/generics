all:
.PHONY: all

test:
	go test -v -race -cover -covermode=atomic ./...
.PHONY: test

bench:
	go test -benchmem -bench ^Benchmark ./...
.PHONY: bench

lint:
	golangci-lint run --fast
.PHONY: lint
