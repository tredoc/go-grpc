package main

import (
	"context"
	"fmt"
	"github.com/tredoc/go-grpc/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Responser struct {
	pb.UnimplementedResponserServer
}

// Ping Unary
func (p *Responser) Ping(_ context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	fmt.Println("Received request:", in.Message)
	time.Sleep(1 * time.Second)
	fmt.Println("Sending pong")
	return &pb.PongResponse{Message: "Pong"}, nil
}

// GetList Server streaming
func (p *Responser) GetList(in *pb.GetListRequest, stream pb.Responser_GetListServer) error {
	for i := range int(in.Count) {
		fmt.Println("Sending number:", i)
		if err := stream.Send(&pb.GetListResponse{Num: int64(i) + 1}); err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

// SendList Client streaming
func (p *Responser) SendList(stream pb.Responser_SendListServer) error {
	return nil
}

// HandleJob Bidirectional streaming
func (p *Responser) HandleJob(stream pb.Responser_HandleJobServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterResponserServer(grpcServer, &Responser{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
