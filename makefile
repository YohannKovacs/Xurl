build:
	@go build -o bin/xurl

run: build
	@./bin/xurl

test:
	@go test -v ./...