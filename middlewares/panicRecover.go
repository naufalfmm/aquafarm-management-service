package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

type PanicError struct {
	Err error
}

func (pe PanicError) Error() string {
	return fmt.Errorf("[PANIC RECOVER] %v", pe.Err).Error()
}

func (m middlewares) PanicRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				r := recover()
				if r != nil {
					err, ok := r.(error)
					if !ok {
						err = &PanicError{
							Err: err,
						}
					}
					generateResp.NewJSONResponse(c, http.StatusInternalServerError, "", err)
				}
			}()
			return next(c)
		}
	}
}
