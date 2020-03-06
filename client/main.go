package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "testProject/proto"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("localhost:8800", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServerClient(conn)
	request := &pb.HelloWorldRequest{
		Text: "Hello, world",
	}
	response, err := client.HelloWorld(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Ans)
}
