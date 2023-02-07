package middlewares

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (m middlewares) RequestStart() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(consts.XRequestIDHeader, uuid.New().String())

			req := dto.RequestStartLogRequest{}
			if err := req.FromEchoContext(c); err != nil {
				return err
			}

			if _, err := m.Usecases.EndpointLogs.RequestStart(context.Background(), req); err != nil {
				return err
			}

			return next(c)
		}
	}
}
