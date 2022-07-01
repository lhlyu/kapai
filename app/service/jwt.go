package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"kapai/app/util"
	"os"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	var err error
	block, _ := pem.Decode([]byte(os.Getenv("JWT_PRIVATE")))
	//X509解码
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		util.Log.Panic().Err(err).Msg("ParsePKCS1PrivateKey")
	}
	block, _ = pem.Decode([]byte(os.Getenv("JWT_PUBLIC")))
	publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		util.Log.Panic().Err(err).Msg("ParsePKCS1PublicKey")
	}
}

// 生成token
func getJwtToken(account string) string {
	tok, err := jwt.NewBuilder().
		Issuer(`kapai`).
		IssuedAt(time.Now()).
		Subject(account).
		Build()
	if err != nil {
		util.Log.Error().Err(err).Str("account", account).Msg("getJwtToken")
		return ""
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		util.Log.Error().Err(err).Str("account", account).Msg("getJwtToken Sign")
		return ""
	}
	return string(signed)
}

// 解析token
func ParseJwtToken(token string) string {
	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, publicKey))
	if err != nil {
		util.Log.Error().Err(err).Str("token", token).Msg("parseJwtToken Sign")
		return ""
	}
	return verifiedToken.Subject()
}
