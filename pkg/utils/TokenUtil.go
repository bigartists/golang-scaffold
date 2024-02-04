package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId int64, privateKey []byte, expireTime time.Time) (string, error) {
	prikey, _ := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	user := UserClaim{UserId: userId}
	user.ExpiresAt = expireTime.Unix()
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, user)
	token, err := tokenObj.SignedString(prikey)
	return token, err
}

func ParseToken(token string, publicKey []byte) (*jwt.Token, error) {
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	getToken, err := jwt.ParseWithClaims(token, &UserClaim{}, func(token *jwt.Token) (i interface{}, e error) {
		return pubKey, nil
	})
	return getToken, err
}
