CLI_NAME=strnaming

default: fmt-code vet-code test bench

build-cli: fmt-code vet-code test
	@rm -rf $(CLI_NAME) &>/dev/null
	@echo "build strnaming cli tool"
	@go build ./cmd/...

fmt-code:
	@echo "fmt code..."
	@go fmt ./...

vet-code:
	@echo "vet code..."
	@go vet ./...

test:
	@echo "run testing..."
	@go test ./...

bench:
	@echo "run benchmark..."
	@go test -bench=. -run=^$
