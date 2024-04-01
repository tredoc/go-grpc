package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/tredoc/go-grpc/proto/gen"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewResponserClient(conn)
	ctx := context.Background()

	// Unary
	{
		fmt.Println("Sending ping")
		resp, err := client.Ping(ctx, &pb.PingRequest{Message: "Ping"})
		if err != nil {
			panic(err)
		}
		fmt.Println("Received response:", resp.Message)
	}
	time.Sleep(5 * time.Second)

	// Server streaming
	{
		count := int64(25)
		fmt.Println("Requesting list of", count, "numbers")
		stream, err := client.GetList(ctx, &pb.GetListRequest{Count: count})
		result := make([]int64, 0, count)
		if err != nil {
			panic(err)
		}
		for {
			resp, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					panic(err)
				}
			}
			fmt.Println("Received number:", resp.Num)
			result = append(result, resp.Num)
		}
		fmt.Println("Received list:", result)
	}
	time.Sleep(5 * time.Second)

	// Client streaming
	// Bidirectional streaming
}
