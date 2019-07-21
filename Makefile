.PHONY: all
all: test benchmark

.PHONY: test
test:
	@go test -v ./...

.PHONY: benchmark
benchmark:
	@go test \
	  -v ./... \
	  -run=^$ \
	  -bench=. \
	  -benchmem \
	  -timeout=1m
