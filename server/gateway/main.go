package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
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
	serverConfig := []struct {
		name       string
		addr       string
		registFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:       "auth",
			addr:       "localhost:8081",
			registFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:       "rental",
			addr:       "localhost:8082",
			registFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}
	for _, s := range serverConfig {
		err := s.registFunc(c, mux, s.addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatalf("cannot register %s service with service: %v", s.name, err)
		}
	}
	log.Fatal(http.ListenAndServe(":8080", mux))

}
