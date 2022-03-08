package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen : %v", zap.Error(err))
	}
	c := context.Background()
	mgURI := "mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(mgURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb: %v", zap.Error(err))
	}

	// 读取私钥
	f, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private.key:%v", zap.Error(err))
	}
	b, err3 := ioutil.ReadAll(f)
	if err3 != nil {
		logger.Fatal("cannot read private.key:%v", zap.Error(err3))
	}
	pk, err4 := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err4 != nil {
		logger.Fatal("cannot Parse private.key:%v", zap.Error(err4))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx0ebc7b6b12ef1585",
			AppSecret: "534a58b2185cdbc3dba6168500b7ad3d",
		},
		Mongo:          dao.NewMongo(mongoClient.Database("coolcar")),
		Logger:         *logger,
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJWTTokenGen("coolcar/auth", pk),
	})
	err2 := s.Serve(lis)
	if err2 != nil {
		logger.Fatal("cannot server: %v", zap.Error(err2))
	}
}
