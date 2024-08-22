package pb

import (
	"fmt"
	"time"

	"github.com/vinit-chauhan/grpc-demo/pb/invoice"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExampleInvoice() {
	// TODO: Change to protobuf invoice
	inv := invoice.Invoice{
		Id:       "2023-0123",
		Time:     timestamppb.New(time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)),
		Customer: "Wile E. Coyote",
		Items: []*invoice.LineItem{
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
	fmt.Printf("%v\n", &inv) // Make compiler happy
	data, err := proto.Marshal(&inv)
	if err == nil {
		fmt.Println("size:", len(data))
	} else {
		fmt.Println("ERROR:", err)
	}

	// Output:
	// &{2023-0123 2023-01-07 13:45:00 +0000 UTC Wile E. Coyote [{hammer-20 1 249} {nail-9 100 1} {glue-5 1 799}]}
	// size: 0
}
