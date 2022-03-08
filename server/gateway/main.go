package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  true,
			}},
	))
	// 监听 auth
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cannot register with service: %v", err)
	}

	// 监听 trip
	err = rentalpb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8082", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cannot register with service: %v", err)
	}
	fmt.Println("start listening trip")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
