all:
.PHONY: all

test:
	go test -cover ./...
.PHONY: test

bench:
	go test -cover -benchmem -bench ^Benchmark ./...
.PHONY: bench

lint:
	golangci-lint run --fast
.PHONY: lint
