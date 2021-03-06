package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math"
	"net"
	pb "testProject/proto"
)

type server struct{}

func main() {

	// init grpc server
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	// pb - package which was generated by *.proto file and protoc command
	pb.RegisterExampleServerServer(grpcServer, &server{})
	grpcServer.Serve(listener)

}

func (s *server) HelloWorld(c context.Context, request *pb.HelloWorldRequest) (*pb.HelloWorldReply, error) {
	fmt.Println(request.Text)
	return &pb.HelloWorldReply{Ans: "successfully printed"}, nil
}

func (s *server) CheckPoints(c context.Context, request *pb.CheckPointsRequest) (*pb.CheckPointsReply, error) {
	ans := checkPoints(request.Rect, request.Points)
	return &pb.CheckPointsReply{IsInRectangle: ans}, nil

}

func checkPoints(rect *pb.Rectangle, points []*pb.Point) map[int32]bool {
	dx := math.Abs(rect.A.X - rect.C.Y)
	minX := math.Min(rect.A.X, rect.C.X)
	dy := math.Abs(rect.A.Y - rect.C.X)
	minY := math.Min(rect.A.Y, rect.C.Y)

	ans := make(map[int32]bool)

	for _, p := range points {
		if math.Abs(p.X-minX) <= dx && math.Abs(p.Y-minY) <= dy {
			ans[p.Id] = true
		} else {
			ans[p.Id] = false
		}
	}
	return ans
}
