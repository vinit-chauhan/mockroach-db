package server

import (
	"context"
	"log"
	"net"

	pb "github.com/vinit-chauhan/grpc-demo/pb/ride"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Rides struct {
	pb.UnimplementedRidesServer
}

func (r Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Println("[Start] function invoked")
	resp := pb.StartResponse{
		Id: req.Id,
	}

	return &resp, nil
}

func (r Rides) End(ctx context.Context, req *pb.EndRequest) (*pb.EndResponse, error) {
	log.Println("[End] function invoked")
	resp := pb.EndResponse{
		Id: req.Id,
	}
	return &resp, nil
}

func Run(ctx context.Context, started chan<- bool) {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	srv := grpc.NewServer()
	var u Rides

	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	log.Printf("server listening on address: %s", addr)

	go func() {
		log.Println("sending signal to start client")
		started <- true
	}()

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
