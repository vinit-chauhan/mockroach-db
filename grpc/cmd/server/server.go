package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/vinit-chauhan/grpc-demo/proto/ride"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Rides struct {
	pb.UnimplementedRidesServer
}

func (r Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Println("[server][Start] function invoked")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "error fetching metadata from context")
	}

	fmt.Println(md.Get("api_key"))

	resp := pb.StartResponse{
		Id: req.Id,
	}

	return &resp, nil
}

func (r Rides) End(ctx context.Context, req *pb.EndRequest) (*pb.EndResponse, error) {
	log.Println("[server][End] function invoked")
	resp := pb.EndResponse{
		Id: req.Id,
	}
	return &resp, nil
}

func (r Rides) Location(stream pb.Rides_LocationServer) error {

	count := int64(0)
	DriverId := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "can't read")
		}

		DriverId = req.DriverId
		count++

		if count == 50 {
			log.Printf("[server][Location] limit reached!!!")
			break
		}
	}

	resp := pb.LocationResponse{
		DriverId: DriverId,
		Count:    count,
	}

	return stream.SendAndClose(&resp)
}

// It's Unary so It will only run for Start and End.
func timingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("[server][timingInterceptor] %s took %v", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}

func Run(ctx context.Context, started chan<- bool) {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	srv := grpc.NewServer(grpc.UnaryInterceptor(timingInterceptor))
	var u Rides

	pb.RegisterRidesServer(srv, &u)
	reflection.Register(srv)

	log.Printf("[server] server listening on address: %s", addr)

	go func() {
		log.Println("[server] sending signal to start client")
		started <- true
	}()

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
