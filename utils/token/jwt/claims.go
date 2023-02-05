package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type UserClaims struct {
	*jwt.StandardClaims
	UserLogin UserLogin `json:"user"`
}

func (c *UserClaims) Valid() error {
	jwt.TimeFunc = time.Now

	return c.StandardClaims.Valid()
}

func (c *UserClaims) SetExp(exp int64) token.Claims {
	c.ExpiresAt = exp

	return c
}
