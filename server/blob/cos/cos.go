package cos

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Service struct {
	client *cos.Client
	secID  string
	secKey string
}

func NewService(addr, secID, seckey string) (*Service, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse addr:%v", err)
	}
	b := &cos.BaseURL{BucketURL: u}
	return &Service{
		client: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secID,
				SecretKey: seckey,
			},
		}),
		secID:  secID,
		secKey: seckey,
	}, nil
}

func (s *Service) SingURL(c context.Context, method, path string, timeout time.Duration) (string, error) {
	u, err := s.client.Object.GetPresignedURL(c, method, path, s.secID, s.secKey, timeout, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
func (s *Service) Get(c context.Context, path string) (io.ReadCloser, error) {
	r, err := s.client.Object.Get(c, path, nil)
	var b io.ReadCloser
	if r != nil {
		b = r.Body
	}
	if err != nil {
		return b, err
	}
	if r.StatusCode > 400 {
		return b, fmt.Errorf("got err response:%+v", err)
	}
	return b, nil
}
