package blob

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/dao"

	"go.uber.org/zap"
)

type service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

func (s *service) CreateBlob(c context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	return nil, nil
}
func (s *service) GetBlob(c context.Context, req *blobpb.GetBlobRequest) (*blobpb.GetBlobResponse, error) {
	return nil, nil
}
func (s *service) GetBlobURL(c context.Context, req *blobpb.GetBlobURlRequest) (*blobpb.GetBlobURLResponse, error) {
	return nil, nil
}
