package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/vinit-chauhan/grpc-demo/pb/ride"
)

func Run(ctx context.Context, start <-chan bool) {
	for {
		select {
		case <-ctx.Done():
			log.Println("[Run] signal from context: closing client")
			return
		case <-start:
			log.Println("[Run] received start signal: starting client")
			run(ctx)
		}
	}

}

func run(ctx context.Context) {
	addr := "localhost:9292"
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[run] error connecting to sever: %s", err)
	}
	defer conn.Close()

	log.Printf("[run] connected to server: %s", addr)

	c := pb.NewRidesClient(conn)

	startReq := pb.StartRequest{
		Id:           "2543750CB50DFSDAF670ASDF70A0",
		DriverId:     "007",
		Location:     &pb.Location{Lat: 38.8951, Lon: -77.0364},
		PassengerIds: []string{"M", "Q"},
		Time:         timestamppb.Now(),
		Type:         pb.RideType_POOL,
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer func() {
		log.Println("[run] timeout: performing cancel()")
		cancel()
	}()

	resp, err := c.Start(ctx, &startReq)
	if err != nil {
		log.Fatalf("[run] error running Start RPC: %s", err)
	}

	log.Printf("[run] response from server: { %v }", resp)
}
