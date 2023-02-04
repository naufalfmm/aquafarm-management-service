package middlewares

import (
	"github.com/labstack/echo/v4"
	echMid "github.com/labstack/echo/v4/middleware"
)

func (m middlewares) RemoveTrailingSlash() echo.MiddlewareFunc {
	return echMid.RemoveTrailingSlash()
}
