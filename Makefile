.PHONY: test-example
test-example:
	go run ./cmd/loglinter/main.go -- ./examples/example.go

GO_UBER_DIR="pkg/analyzer/testdata/src/go.uber.org"

.PHONY: test-deps
test-deps:
	mkdir -p ${GO_UBER_DIR}
	rm -rf ${GO_UBER_DIR}/*
	cd ${GO_UBER_DIR} && git clone https://github.com/uber-go/zap.git
	cd ${GO_UBER_DIR} && git clone https://github.com/uber-go/multierr.git

.PHONY: test
test: test-deps
	go test ./...

.PHONY: build-deps
build-deps:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.9.0

.PHONY: build
build: build-deps
	golangci-lint custom

.PHONY: test-built
test-built: build
	./custom-gcl run examples/example.go
