run/client:
	@echo "Running client"
	@go run cmd/client/main.go

run/server:
	@echo "Running server"
	@go run cmd/server/main.go

run/all:
	@echo "Running all services"
	@go run cmd/server/main.go &
	@go run cmd/client/main.go

gen:
	@echo "Generating proto files"
	@protoc -I proto src/ping.proto \
		--go_out=. \
    	--go-grpc_out=.

.PHONY: run/client run/server run/all gen
.SILENT: run/client run/server run/all gen