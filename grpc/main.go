package main

import (
	"context"
	"sync"

	"github.com/vinit-chauhan/grpc-demo/cmd/client"
	"github.com/vinit-chauhan/grpc-demo/cmd/server"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := make(chan bool)
	wg.Add(1)
	go client.Run(ctx, start)

	server.Run(ctx, start)
	wg.Wait()
}
