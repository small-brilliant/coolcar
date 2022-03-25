package blob

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/dao"
	"coolcar/shared/id"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage interface {
	SingURL(c context.Context, method, path string, timeout time.Duration) (string, error)
	Get(c context.Context, path string) (io.ReadCloser, error)
}
type Service struct {
	Storage Storage
	Mongo   *dao.Mongo
	Logger  *zap.Logger
}

func (s *Service) CreateBlob(c context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	aid := id.AccountID(req.AccoundId)
	br, err := s.Mongo.CreateBlob(c, aid)
	if err != nil {
		s.Logger.Error("cannot create blob:%v", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	u, err := s.Storage.SingURL(c, http.MethodPut, br.Path, secToDuration(req.UploadUrlTimeSec))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot sign url:%v", err)
	}
	return &blobpb.CreateBlobResponse{
		Id:        br.ID.Hex(),
		UploadUrl: u,
	}, nil
}

func (s *Service) GetBlob(c context.Context, req *blobpb.GetBlobRequest) (*blobpb.GetBlobResponse, error) {
	br, err := s.getBlobRecord(c, id.BlobID(req.Id))
	if err != nil {
		return nil, err
	}
	r, err := s.Storage.Get(c, br.Path)
	if r != nil {
		defer r.Close()
	}
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot get Storage:%v", err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot read from response:%v", err)
	}
	return &blobpb.GetBlobResponse{
		Data: b,
	}, nil
}

func (s *Service) GetBlobURL(c context.Context, req *blobpb.GetBlobURlRequest) (*blobpb.GetBlobURLResponse, error) {
	br, err := s.getBlobRecord(c, id.BlobID(req.Id))
	if err != nil {
		return nil, err
	}
	u, err := s.Storage.SingURL(c, http.MethodGet, br.Path, secToDuration(req.TimeSec))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot get Storage:%v", err)
	}
	return &blobpb.GetBlobURLResponse{
		Url: u,
	}, nil
}

func (s *Service) getBlobRecord(c context.Context, bid id.BlobID) (*dao.BlobRecord, error) {
	br, err := s.Mongo.GetBlob(c, bid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "")
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return br, nil
}

func secToDuration(sec int32) time.Duration {
	return time.Duration(sec) * time.Second
}
