package generateResp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
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

func NewJSONResponse(ec echo.Context, statusCode int, message string, data interface{}) error {
	if statusCode >= http.StatusBadRequest {
		resp := Error{
			Default: Default{
				Ok:      false,
				Message: message,
			},
			Error: data.(error).Error(),
		}

		ec.Set(consts.ApiResponse, resp)

		return ec.JSON(statusCode, resp)
	}

	resp := Success{
		Default: Default{
			Ok:      true,
			Message: "Success",
		},
		Data: data,
	}

	ec.Set(consts.ApiResponse, resp)

	return ec.JSON(statusCode, resp)
}
