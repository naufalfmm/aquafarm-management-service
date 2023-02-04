package generateResp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Default struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}

	Error struct {
		Default
		Error string `json:"error"`
	}

	Success struct {
		Default
		Data interface{} `json:"data,omitempty"`
	}
)

func NewJSONResponse(eCtx echo.Context, statusCode int, message string, data interface{}) error {
	if statusCode >= http.StatusBadRequest {
		return eCtx.JSON(statusCode, Error{
			Default: Default{
				Ok:      false,
				Message: message,
			},
			Error: data.(error).Error(),
		})
	}

	return eCtx.JSON(statusCode, Success{
		Default: Default{
			Ok:      true,
			Message: "Success",
		},
		Data: data,
	})
}
