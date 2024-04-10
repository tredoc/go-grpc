package main

import (
	"context"
	"fmt"
	pb "github.com/tredoc/go-grpc/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net"
	"time"
)

type Responser struct {
	pb.UnimplementedResponserServer
}

// Ping Unary
func (p *Responser) Ping(_ context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	fmt.Println("\nReceived request:", in.Message)
	time.Sleep(1 * time.Second)
	fmt.Println("Sending pong")
	return &pb.PongResponse{Message: "Pong"}, nil
}

// GetList Server streaming
func (p *Responser) GetList(in *pb.GetListRequest, stream pb.Responser_GetListServer) error {
	fmt.Println("\nReceived request for list of", in.Count, "numbers")
	for i := range int(in.Count) {
		num := int64(i + 1)
		fmt.Println("Sending number:", num)
		if err := stream.Send(&pb.GetListResponse{Num: num}); err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

// SendList Client streaming
func (p *Responser) SendList(stream pb.Responser_SendListServer) error {
	fmt.Println("\nReceiving stream of timestamps")
	start := time.Now()
	res := make([]*timestamppb.Timestamp, 0, 100)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			err = stream.SendAndClose(&pb.SendListResponse{Timestamp: res})
			if err != nil {
				fmt.Println(err)
			}
			break
		}
		if err != nil {
			return err
		}

		fmt.Println("Received timestamp by server:", in.Timestamp.Seconds)
		res = append(res, in.Timestamp)
	}

	fmt.Println("Handled client streaming in:", time.Since(start).Seconds(), "seconds")
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
