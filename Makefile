
# Lets "make" always run test targets
.PHONY: test 
	
build:
	 @go build -o game-server ./cmd/
	
run: build
	@./bin/game-server

test:
	@go test ./...

