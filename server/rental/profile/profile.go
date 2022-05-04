package profile

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	BlobClient        blobpb.BlobserviceClient
	PhotoGetExpire    time.Duration
	PhotoUploadExpire time.Duration
	Mongo             *dao.Mongo
	Logger            *zap.Logger
}

func (s *Service) GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	p, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		code := s.logAndConvertProfileErr(err)
		if code == codes.NotFound {
			return &rentalpb.Profile{}, nil
		}
		return nil, status.Error(code, "")
	}
	if p.Profile == nil {
		return &rentalpb.Profile{}, nil
	}
	return p.Profile, nil
}
func (s *Service) SubmitProfile(c context.Context, req *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       req,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_UNSUBMITTED, p)
	if err != nil {
		s.Logger.Error("cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	// 模拟对驾驶照喝信息的审核

	// 后台任务
	go func() {
		time.Sleep(3 * time.Second)
		err := s.Mongo.UpdateProfile(context.Background(), aid, rentalpb.IdentityStatus_PENDING, &rentalpb.Profile{
			Identity:       req,
			IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
		})
		if err != nil {
			s.Logger.Error("cannot verigfy identity", zap.Error(err))
		}
	}()
	return p, nil
}
func (s *Service) ClearProfile(c context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	p := &rentalpb.Profile{}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &rentalpb.Profile{}, nil
		}
		s.Logger.Error("cannot get Profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return p, nil
}

func (s *Service) GetProfilePhoto(c context.Context, req *rentalpb.GetProfilePhotoRequest) (*rentalpb.GetProfilePhotoReponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		code := s.logAndConvertProfileErr(err)
		if code == codes.NotFound {
			return &rentalpb.GetProfilePhotoReponse{}, nil
		}
		return nil, status.Error(code, "")
	}
	if pr.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}
	br, err := s.BlobClient.GetBlobURL(c, &blobpb.GetBlobURlRequest{
		Id:      pr.PhotoBlobID,
		TimeSec: int32(s.PhotoGetExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("cannot get blob", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &rentalpb.GetProfilePhotoReponse{
		Url: br.Url,
	}, nil
}
func (s *Service) CreateProfilePhoto(c context.Context, req *rentalpb.CreateProfilePhotoRequest) (*rentalpb.CreateProfilePhotoReponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	br, err := s.BlobClient.CreateBlob(c, &blobpb.CreateBlobRequest{
		AccoundId:        aid.String(),
		UploadUrlTimeSec: int32(s.PhotoUploadExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("cannot create blob", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}
	err = s.Mongo.UpdateProfilePhoto(c, aid, id.BlobID(br.Id))
	if err != nil {
		s.Logger.Error("cannot update blob", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}
	return &rentalpb.CreateProfilePhotoReponse{
		UploadUrl: br.UploadUrl,
	}, nil
}
func (s *Service) CompleteProfilePhoto(c context.Context, req *rentalpb.CompleteProfilePhotoRequest) (*rentalpb.Identity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		return nil, status.Error(s.logAndConvertProfileErr(err), "")
	}
	if pr.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}
	br, err := s.BlobClient.GetBlob(c, &blobpb.GetBlobRequest{
		Id: pr.PhotoBlobID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	s.Logger.Info("got profile photo", zap.Int("size", len(br.Data)))
	return &rentalpb.Identity{
		LicNumber:       "34088119923324908",
		Name:            "谭谭",
		Gender:          rentalpb.Gender_FEMALE,
		BirthDateMillis: 23498237948237,
	}, nil
}

func (s *Service) ClearProfilePhoto(c context.Context, req *rentalpb.ClearProfilePhotoRequest) (*rentalpb.ClearProfilePhotoReponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	err = s.Mongo.UpdateProfilePhoto(c, aid, id.BlobID(""))
	if err != nil {
		s.Logger.Error("cannot clear profile photo", zap.Error(err))
	}
	return &rentalpb.ClearProfilePhotoReponse{}, nil
}

func (s *Service) logAndConvertProfileErr(err error) codes.Code {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return codes.NotFound
		}
		s.Logger.Error("cannot get Profile", zap.Error(err))
		return codes.Internal
	}
	return codes.OK
}
