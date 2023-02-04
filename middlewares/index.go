package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/resources"
)

type (
	Middlewares interface {
		RemoveTrailingSlash() echo.MiddlewareFunc
		PanicRecover() echo.MiddlewareFunc
		ImplementCors() echo.MiddlewareFunc
	}

	middlewares struct {
		Resource resources.Resources
	}
)

func Init(resources resources.Resources) (Middlewares, error) {
	return &middlewares{resources}, nil
}
