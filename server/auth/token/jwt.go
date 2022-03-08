package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGen struct {
	PrivateKey *rsa.PrivateKey
	Issuer     string
	// 希望时间可以又外面控制，这样测试的时候不会因为时间导致不一样
	nowFun func() time.Time
}

// 构造函数

func NewJWTTokenGen(issue string, privatekey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		Issuer:     issue,
		nowFun:     time.Now,
		PrivateKey: privatekey,
	}
}

func (t *JWTTokenGen) GeneratorToken(accountID string, expire time.Duration) (string, error) {
	nowSec := t.nowFun().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.Issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(expire.Seconds()),
		Subject:   accountID,
	})
	//签名
	return tkn.SignedString(t.PrivateKey)
}
