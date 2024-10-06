
# Lets "make" always run test targets
.PHONY: test 
	
build:
	 @go build -o bin/game-server ./cmd/
	
run: build
	@./bin/game-server

test:
	@go test ./...



	
# build:
# 	 @go build -o bin/starlight-cargo ./cmd/app/
# 	
# run: build
# 	@./bin/starlight-cargo
#
