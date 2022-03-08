package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	//1.解析token，结构得对
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		//验证签名
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}
	// 验证token是由有效
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*jwt.StandardClaims)
	//保证是StandardClaim
	if !ok {
		return "", fmt.Errorf("token is not StandardClaim")
	}
	//检查过期时间等等
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not vaild:%v", err)
	}
	return clm.Subject, nil
}
