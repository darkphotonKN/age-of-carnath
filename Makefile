.PHONY: test 
	
build:
	 @go build -o bin/game-server ./cmd/
	
run: build
	@./bin/game-server

test:
	@go test ./... --cover



