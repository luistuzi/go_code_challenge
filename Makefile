build:
	@go build -o bin/go_code_challenge main.go

test:
	@test -v./...

run: build
	@./bin/go_code_challenge
	