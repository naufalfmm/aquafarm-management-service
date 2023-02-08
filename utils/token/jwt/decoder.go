package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type decoder struct {
	publicKey string
}

func NewDecoder(publicKey string) token.Decoder {
	return &decoder{
		publicKey: publicKey,
	}
}

func (d *decoder) DecodeToken(t string) (token.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(t, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(d.publicKey), nil
	})
	if err != nil || tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, consts.ErrUnclaimedToken
	}

	tokClaims := tokenClaims.Claims.(*UserClaims)
	return tokClaims, nil
}
