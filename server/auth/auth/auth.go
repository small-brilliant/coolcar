package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth/dao"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         zap.Logger
}
type OpenIDResolver interface {
	Resolver(code string) (string, error)
}
type TokenGenerator interface {
	GeneratorToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("recived code", zap.String("code", req.Code))
	openId, err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		//这个错误可能包含了内部服务的信息
		return nil, status.Errorf(codes.Unavailable, "cannot resolve opendId : %v", err)
	}
	accountID, err := s.Mongo.ResolverAccountID(c, openId)
	if err != nil {
		s.Logger.Error("cannot resolve account id : %v", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "")
	}
	tkn, err := s.TokenGenerator.GeneratorToken(accountID.String(), 2*time.Hour)
	if err != nil {
		s.Logger.Error("cannot generator token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresInt:  int32(s.TokenExpire.Seconds()),
	}, nil
}
