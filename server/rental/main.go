package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile"
	profileDao "coolcar/rental/profile/dao"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	profileClient "coolcar/rental/trip/client/profile"
	tripDao "coolcar/rental/trip/dao"

	"coolcar/shared/server"
	"log"

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
	c := context.Background()
	mgURI := "mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(mgURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb: %v", zap.Error(err))
	}
	db := mongoClient.Database("coolcar")
	profileService := &profile.Service{
		Mongo:  profileDao.NewMongo(db),
		Logger: logger,
	}
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		Logger:            logger,
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{},
				ProfileManager: &profileClient.Manager{
					Fetcher: profileService,
				},
				POIManager: &poi.Manager{},
				Mongo:      tripDao.NewMongo(db),
				Logger:     logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profileService)
		},
	})
	if err != nil {
		logger.Fatal("cannot server: %v", zap.Error(err))
	}
}
