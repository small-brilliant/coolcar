package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(l))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	// 调用cancel内部就和gateway断开
	defer cancel()
	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		//注册位置
		mux,
		":8081",
		//连接方式
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannot start gateway: %v", err)
	}

	err2 := http.ListenAndServe(":8080", mux)
	if err2 != nil {
		log.Fatalf("err2:%v", err2)
	}
}
