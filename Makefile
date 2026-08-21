build:
	@go build -o bin/go_code_challenge cmd/main.go

test:
	@test -v./...

run: build
	@./bin/go_code_challenge
