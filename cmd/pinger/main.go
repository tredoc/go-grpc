package main

import (
	"context"
	"fmt"
	pb "github.com/tredoc/go-grpc/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPongClient(conn)

	fmt.Println("Sending ping")
	resp, err := client.Ping(context.Background(), &pb.PingRequest{Message: "Ping"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Received response:", resp.Message)
}
