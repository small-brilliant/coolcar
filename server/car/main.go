package main

import (
	"context"
	"coolcar/car/amqpclient"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/car"
	"coolcar/car/dao"
	"coolcar/car/sim"
	"coolcar/shared/server"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger : %v", err)
	}
	//1.获取GRPC服务
	c := context.Background()
	mgURI := "mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(mgURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb: %v", zap.Error(err))
	}
	db := mongoClient.Database("coolcar")

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Fatal("cannot dial amqp", zap.Error(err))
	}
	pub, err := amqpclient.NewPublisher(amqpConn, "coolcar")
	if err != nil {
		logger.Fatal("cannot create publisher", zap.Error(err))
	}

	// 后台模拟
	// 自己连自己，grpc.Dial在发请求的时候连接
	carConn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("cannot connection carservice", zap.Error(err))
	}
	sub, err := amqpclient.NewSubscriber(amqpConn, "coolcar", logger)
	if err != nil {
		logger.Fatal("cannot create subscriber", zap.Error(err))
	}
	simController := &sim.Controller{
		CarService: carpb.NewCarServiceClient(carConn),
		Subscriber: sub,
		Logger:     logger,
	}
	go simController.RunSimulations(context.Background())

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:   "car",
		Addr:   ":8084",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			carpb.RegisterCarServiceServer(s, &car.Service{
				Mongo:     dao.NewMongo(db),
				Logger:    logger,
				Publisher: pub,
			})
		},
	})
	if err != nil {
		logger.Fatal("cannot server: %v", zap.Error(err))
	}
}
