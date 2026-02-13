.PHONY: test-example
test-example:
	go run ./cmd/loglinter/main.go -- ./examples/example.go
