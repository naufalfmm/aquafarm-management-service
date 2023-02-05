package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type encoder struct {
	privateKey string
	alg        string
}

func NewEncoder(privateKey, alg string) token.Encoder {
	return &encoder{
		privateKey: privateKey,
		alg:        alg,
	}
}

func (e *encoder) EncodeToken(claims token.Claims) (string, error) {
	privateKey := []byte(e.privateKey)

	newToken := jwt.New(jwt.GetSigningMethod(e.alg))
	newToken.Claims = claims

	signedToken, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
