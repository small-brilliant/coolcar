package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"log"
	"coolcar/blob/blob"
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
	mgURI := "mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mgURI))
	db := mongoClient.Database("coolcar")
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:   "blob",
		Addr:   "8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobserviceServer(s, &blob.service{
				Mongo:  dao.NewMongo(db),
				Logger: logger,
			})
		},
	})
}
