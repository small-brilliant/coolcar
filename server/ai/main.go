package main

import (
	aipb "coolcar/ai/api/gen/v1"
	"coolcar/car/mq/amqpclient"
	"coolcar/shared/server"
	"log"

	"coolcar/ai/ai"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger : %v", err)
	}
	// mq
	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Fatal("cannot dial amqp", zap.Error(err))
	}
	pub, err := amqpclient.NewPublisher(amqpConn, "pos_sim")
	if err != nil {
		logger.Fatal("cannot create publisher", zap.Error(err))
	}
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:   "ai",
		Addr:   ":18001",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			aipb.RegisterAiServiceServer(s, &ai.Service{
				Logger:    logger,
				Publisher: pub,
			})
		},
	})
	if err != nil {
		logger.Fatal("cannot server: %v", zap.Error(err))
	}
}
