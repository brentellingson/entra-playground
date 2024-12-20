.PHONY: all test lint format generate restore

test: format
	go test ./... --cover -coverprofile=coverage.out -covermode=count -v

lint: format
	golangci-lint run

format: generate
	swag fmt
	gci write -s standard -s default .
	gofumpt -l -w .

generate: restore
	swag init --generalInfo ./internal/api/server.go --output ./internal/docs

restore:
	go get -v ./...
	go mod tidy
