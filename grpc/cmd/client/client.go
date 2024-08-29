package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vinit-chauhan/grpc-demo/proto/ride"
)

func Run(ctx context.Context, start <-chan bool) {
	for {
		select {
		case <-ctx.Done():
			log.Println("[client][Run] signal from context: closing client")
			return
		case <-start:
			log.Println("[client][Run] received start signal: starting client")
			run(ctx)
		}
	}

}

const (
	Id       = "2543750CB50DFSDAF670ASDF70A0"
	DriverId = "007"
)

func run(ctx context.Context) {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[client][run] error connecting to sever: %s", err)
	}
	defer conn.Close()

	log.Printf("[client][run] connected to server: %s", addr)

	c := pb.NewRidesClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer func() {
		log.Println("[client][run] timeout: performing cancel()")
		cancel()
	}()

	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "testK3y")

	start(ctx, c)
	end(ctx, c)
	stream(ctx, c)

}

func stream(ctx context.Context, c pb.RidesClient) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer func() {
		log.Println("[client][stream] timeout: performing cancel()")
		cancel()
	}()

	fmt.Println(Id)

	stream, err := c.Location(ctx)
	if err != nil {
		log.Fatalf("[client][stream] error executing Location RPC: %s", err)
	}

	lreq := pb.LocationRequest{
		DriverId: DriverId,
		Location: &pb.Location{Lat: 38.8951, Lon: -77.0354},
	}

	stream.Send(&lreq)

	for range 100 {
		lreq.Location.Lon -= 0.001
		if err := stream.Send(&lreq); err != nil {
			log.Fatalf("[client][stream] error sending location: %s", err)
			break
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[client][stream] error executing Location RPC: %s", err)
	}
	log.Printf("[client][stream] Location response from server: { %v }", resp)
}

func end(ctx context.Context, c pb.RidesClient) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer func() {
		log.Println("[client][end] timeout: performing cancel()")
		cancel()
	}()

	req := pb.EndRequest{
		Id:       Id,
		EndTime:  timestamppb.Now(),
		Distance: 1.4,
	}

	resp, err := c.End(ctx, &req)
	if err != nil {
		log.Fatalf("[client][end] error executing End RPC: %s", err)
	}
	log.Printf("[client][end] End response from server: { %v }", resp)
}

func start(ctx context.Context, c pb.RidesClient) {

	req := pb.StartRequest{
		Id:           Id,
		DriverId:     DriverId,
		Location:     &pb.Location{Lat: 38.8951, Lon: -77.0364},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer func() {
		log.Println("[client][start] timeout: performing cancel()")
		cancel()
	}()

	resp, err := c.Start(ctx, &req)
	if err != nil {
		log.Fatalf("[client][start] error executing Start RPC: %s", err)
	}
	log.Printf("[client][start] Start response from server: { %v }", resp)

}
