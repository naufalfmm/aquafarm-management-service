package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
)

func (m middlewares) VerifyToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenAuth := c.Request().Header.Get("Authorization")
			if tokenAuth == "" {
				return generateResp.NewJSONResponse(c, http.StatusUnauthorized, "", consts.ErrInvalidToken)
			}

			token := ""
			tokenSplit := strings.Split(tokenAuth, " ")
			if len(tokenSplit) > 1 {
				token = tokenSplit[1]
			}

			tokenData, err := m.Resources.JWT.Decoder.DecodeToken(token)
			if err != nil {
				return generateResp.NewJSONResponse(c, http.StatusUnauthorized, "", consts.ErrInvalidToken)
			}

			userTokenData := tokenData.(*jwt.UserClaims).UserLogin

			c.Set(consts.XUserHeader, &userTokenData)

			return next(c)
		}
	}
}
