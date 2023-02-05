package jwToken

import (
	"github.com/naufalfmm/aquafarm-management-service/resources/config"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
)

func NewJWT(config *config.EnvConfig) (jwt.JWT, error) {
	return jwt.NewJWT(config.JwtPublicKey, config.JwtAlg)
}
