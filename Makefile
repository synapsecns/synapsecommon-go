

.PHONY: test lint fmt fmt-all generate test-local setup-local

test:
	@go test ./... -v

test-ci:
	go test -v -coverprofile=profile.cov ./...

test-local: setup-local
	@./scripts/shdotenv -e .env go test -v ./...

lint:
	@golangci-lint run --config .golangci.yml

fmt:
	@gofmt -s -w $(FILE)

fmt-all:
	@gofmt -s -w ./..

setup-local:
	@mkdir -p ./scripts
	@if [[ ! -f ./scripts/shdotenv ]]; then wget https://github.com/ko1nksm/shdotenv/releases/latest/download/shdotenv -O ./scripts/shdotenv; chmod +x ./scripts/shdotenv; fi