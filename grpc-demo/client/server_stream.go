package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vinit-chauhan/grpc-demo/proto"
)

func callSatHelloServerSideStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")

	stream, err := client.SatHelloServerSideStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}

		log.Println(msg)
	}

	log.Println("Streaming finished")
}
