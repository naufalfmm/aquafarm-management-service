package middlewares

import (
	"github.com/labstack/echo/v4"
	echMid "github.com/labstack/echo/v4/middleware"
)

func (m middlewares) ImplementCors() echo.MiddlewareFunc {
	return echMid.CORSWithConfig(echMid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	})
}
