package main

import (
	"context"
	"fmt"
	"github.com/tredoc/go-grpc/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Ponger struct {
	pb.UnimplementedPongServer
}

func (p *Ponger) Ping(_ context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	fmt.Println("Received request:", in.Message)
	fmt.Println("Sending pong")
	return &pb.PongResponse{Message: "Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPongServer(grpcServer, &Ponger{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
