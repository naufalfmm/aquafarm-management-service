package middlewares

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
)

func (m middlewares) RequestStart() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(consts.XRequestIDHeader, uuid.New().String())
			c.Set(consts.XRequestStartUnix, time.Now().UnixMilli())

			return next(c)
		}
	}
}
