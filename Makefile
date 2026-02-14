.PHONY: test-example
test-example:
	go run ./cmd/loglinter/main.go -- ./examples/example.go

GO_UBER_DIR="pkg/analyzer/testdata/src/go.uber.org"

.PHONY: test-deps
test-deps:
	mkdir -p ${GO_UBER_DIR}
	cd ${GO_UBER_DIR} && git clone https://github.com/uber-go/zap.git
	cd ${GO_UBER_DIR} && git clone https://github.com/uber-go/multierr.git

.PHONY: test
test:
	go test ./...
