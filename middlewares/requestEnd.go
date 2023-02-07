package middlewares

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (m middlewares) RequestEnd() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				_, ok := c.Get(consts.XRequestIDHeader).(string)
				if !ok {
					return
				}

				req := dto.RecordRequestLogRequest{}
				if err := req.FromEchoContext(c); err != nil {
					return
				}

				go m.Usecases.EndpointLogs.RecordRequestLog(context.Background(), req)
			}()
			return next(c)
		}
	}
}
