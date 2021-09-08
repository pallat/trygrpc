package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "poc/server/helloworldpb"

	"google.golang.org/grpc"
)

const port = 9092

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	fmt.Println("serve...")
	grpcServer.Serve(lis)
}

type greeterServer struct{ pb.GreeterServer }

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("greeting to", req.GetName())

	return &pb.HelloReply{
		Message: "this is our first GRPC",
	}, nil
}
func (s *greeterServer) SayHelloAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("greeting to", req.GetName())

	return &pb.HelloReply{
		Message: "this is our GRPC again",
	}, nil
}
