package main

import (
	"log"
	"net"

	pb "github.com/vinit-chauhan/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	host = "localhost:8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server %v", err)
	}

}
