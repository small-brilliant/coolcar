package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ImpersonateAccoutHeader defines the head for accountid
const ImpersonateAccoutHeader = "impersonate-account-id"

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file: %v", err)
	}
	pk, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key : %v", err)
	}

	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pk,
		},
	}
	return i.HandReq, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}
type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	aid := impersonateFromContext(ctx)
	if aid != "" {
		fmt.Printf("impersonating %v\n", aid)
		return handler(ContestWithAccontId(ctx, id.AccountID(aid)), req)
	}
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err = i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}
	return handler(ContestWithAccontId(ctx, id.AccountID(aid)), req)
}
func impersonateFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}
	imp := m[ImpersonateAccoutHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}
func tokenFromContext(c context.Context) (string, error) {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	tkn := ""
	for _, v := range m["authorization"] {
		if strings.HasPrefix(v, "Bearer ") {
			tkn = v[len("Bearer "):]
		}
	}
	if tkn == "" {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return tkn, nil
}

type accountIDKey struct{}

func ContestWithAccontId(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func AccountIDFromContext(c context.Context) (id.AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(id.AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
