.PHONY: all
all: test benchmark

.PHONY: test
test:
	@\
	go test -v ./...

.PHONY: benchmark
benchmark:
	@\
	GOMAXPROCS=1 \
	\
	go test \
		-v ./... \
		-run=^@$ \
		-bench=. \
		-benchmem \
		-benchtime=5s
