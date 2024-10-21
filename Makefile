.PHONY: test 
	
build:
	 @go build -o bin/game-server ./cmd/
	
run: build
	@./bin/game-server

test:
	@go test ./... --cover

test-game-preview:
	@go test ./internal/game/ --cover -coverprofile=coverage.out 
	@go tool cover -html=coverage.out


