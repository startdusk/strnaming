CLI_NAME=strnaming

all: fmt vet test testrace

build-cli: all
	@rm -rf $(CLI_NAME) &>/dev/null
	@echo "build strnaming cli tool"
	@go build ./cmd/...

fmt:
	@echo "fmt code..."
	@go fmt ./...

vet:
	@echo "vet code..."
	@go vet ./...

test:
	@echo "run testing..."
	@go test ./...

bench:
	@echo "run benchmark..."
	@go test -bench=. -run=^$

testrace:
	@echo "run test race..."
	@go test -race -cpu 1,4 -timeout 7m ./...

.PHONY: \
		fmt \
		vet \
		test \
		bench \
		testrace \
		build-cli
