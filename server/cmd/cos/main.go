package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := blobpb.NewBlobserviceClient(conn)

	ctx := context.Background()

	// res, err := c.CreateBlob(ctx, &blobpb.CreateBlobRequest{
	// 	AccoundId:        "account_1",
	// 	UploadUrlTimeSec: 1000,
	// })
	// res, err := c.GetBlob(ctx, &blobpb.GetBlobRequest{
	// 	Id: "624feb44ccff44db122fec90",
	// })
	res, err := c.GetBlobURL(ctx, &blobpb.GetBlobURlRequest{
		Id:      "624feb44ccff44db122fec90",
		TimeSec: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
