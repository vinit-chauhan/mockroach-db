package main

import (
	"log"
	"time"

	pb "github.com/vinit-chauhan/grpc-demo/proto"
)

func (s *helloServer) SatHelloServerSideStreaming(req *pb.NamesList, stream pb.GreetService_SatHelloServerSideStreamingServer) error {
	log.Printf("got request with names: %v", req.Name)

	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
