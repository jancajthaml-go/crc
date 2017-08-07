.PHONY: all
all: build benchmark

.PHONY: build
build:
	@go build crc.go main.go

.PHONY: test
test:
	@go test

.PHONY: benchmark
benchmark:
	@go test -bench=.