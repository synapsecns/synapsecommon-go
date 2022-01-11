

.PHONY: test lint fmt fmt-all generate

test:
	@env ROOT_DIR="$(shell echo $$PWD)" go test ./... -v

lint:
	golangci-lint run --config .golangci.yml

fmt:
	gofmt -s -w $(FILE)

fmt-all:
	gofmt -s -w ./..