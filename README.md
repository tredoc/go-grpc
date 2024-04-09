# go-grpc
This is a dummy gRPC server and client implementation in Golang.

### Hint how to install protoc
* `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
* `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
* `PB_REL="https://github.com/protocolbuffers/protobuf/releases"`  
* `curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip`  
* `unzip protoc-3.26.0-linux-x86_64.zip -d protoc3`  
* `sudo mv protoc3/bin/* /usr/local/bin/`  
* `sudo mv protoc3/include/* /usr/local/include/`  

## Flow
[X] Implement unary request/response  
[X] Implement server streaming  
[X] Implement client streaming  
[ ] Implement bi-directional streaming  