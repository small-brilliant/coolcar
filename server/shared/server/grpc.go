package server

import (
	"coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen : %v", nameField, zap.Error(err))
	}
	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor("shared/auth/public.key")
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameField)
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	c.Logger.Info("server stated", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
