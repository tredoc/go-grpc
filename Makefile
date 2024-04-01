run/pinger:
	@echo "Running pinger"
	@go run cmd/pinger/main.go

run/ponger:
	@echo "Running ponger"
	@go run cmd/ponger/main.go

gen:
	@echo "Generating proto files"
	@protoc -I proto src/ping.proto \
		--go_out=. \
         --go-grpc_out=.

.PHONY: run/pinger run/ponger gen
.SILENT: run/pinger run/ponger gen