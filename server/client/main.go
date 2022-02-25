package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}
	tsc := trippb.NewTripServiceClient(cc)
	r, err := tsc.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip423",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip:%v", err)
	}
	fmt.Println(r)
}
