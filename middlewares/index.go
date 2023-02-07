package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
)

type (
	Middlewares interface {
		RemoveTrailingSlash() echo.MiddlewareFunc
		PanicRecover() echo.MiddlewareFunc
		ImplementCors() echo.MiddlewareFunc
		VerifyToken() echo.MiddlewareFunc
		RequestStart() echo.MiddlewareFunc
		RequestEnd() echo.MiddlewareFunc
	}

	middlewares struct {
		Resources resources.Resources
		Usecases  usecases.Usecases
	}
)

func Init(resources resources.Resources, usecases usecases.Usecases) (Middlewares, error) {
	return &middlewares{resources, usecases}, nil
}
