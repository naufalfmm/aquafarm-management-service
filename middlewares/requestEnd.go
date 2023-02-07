package middlewares

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
)

func (m middlewares) RequestEnd() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				requestID, ok := c.Get(consts.XRequestIDHeader).(string)
				if !ok {
					return
				}

				if _, err := m.Usecases.EndpointLogs.RequestEnd(context.Background(), requestID); err != nil {
					return
				}
			}()
			return next(c)
		}
	}
}
