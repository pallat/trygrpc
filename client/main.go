package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "client/helloworldpb"
)

const serverAddr = "127.0.0.1:9092"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	}
	defer conn.Close()

	if conn == nil {
		log.Fatal("connnection is nil")
	}

	client := pb.NewGreeterClient(conn)

	res, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Pallat",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("greeting response is", res.GetMessage())

	{
		res, err := client.SayHelloAgain(context.Background(), &pb.HelloRequest{
			Name: "Gopher",
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("greeting response is", res.GetMessage())
	}
}
