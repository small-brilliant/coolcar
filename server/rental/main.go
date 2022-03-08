package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/auth"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger : %v", err)
	}
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen : %v", zap.Error(err))
	}
	in, err := auth.Interceptor("shared/auth/public.key")
	if err != nil {
		logger.Fatal("cannot create auth interceptor")
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})
	err2 := s.Serve(lis)
	if err2 != nil {
		logger.Fatal("cannot server: %v", zap.Error(err2))
	}
}
