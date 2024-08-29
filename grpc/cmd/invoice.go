package cmd

import (
	"fmt"
	"time"

	pb "github.com/vinit-chauhan/grpc-demo/proto/invoice"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExampleInvoice() {
	inv := pb.Invoice{
		Id:       "2023-0123",
		Time:     timestamppb.New(time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)),
		Customer: "Wile E. Coyote",
		Items: []*pb.LineItem{
			{
				Sku:    "hammer-20",
				Amount: 1,
				Price:  249,
			}, {
				Sku:    "nail-9",
				Amount: 100,
				Price:  1,
			}, {
				Sku:    "glue-5",
				Amount: 1,
				Price:  799,
			},
		},
	}

	data, err := proto.Marshal(&inv)
	if err == nil {
		fmt.Println("size:", len(data))
	} else {
		fmt.Println("ERROR:", err)
	}
}
