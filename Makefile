build-server:
	@go build -o ./build/serverbin ./server/cmd/server/main.go

run-server: build-server
	@./build/serverbin
