package generateResp

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/stretchr/testify/assert"
)

func TestNewJSONResponse(t *testing.T) {
	var (
		anyData = map[string]interface{}{
			"any": "data",
		}
	)

	t.Run("If the status code 2xx, it will return the success format", func(t *testing.T) {
		message := "Any Message"
		statusCode := http.StatusOK

		resp := Success{
			Default: Default{
				Ok:      true,
				Message: "Success",
			},
			Data: anyData,
		}

		expRes, _ := json.Marshal(resp)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		eCtx := echo.New().NewContext(req, res)

		err := NewJSONResponse(eCtx, statusCode, message, anyData)

		assert.Nil(t, err)
		assert.Equal(t, statusCode, res.Code)
		assert.Equal(t, string(expRes)+"\n", res.Body.String())
		assert.NotNil(t, eCtx.Get(consts.ApiResponse))
	})

	t.Run("If the status code 3xx, it will return the success format", func(t *testing.T) {
		message := "Any Message"
		statusCode := http.StatusFound

		resp := Success{
			Default: Default{
				Ok:      true,
				Message: "Success",
			},
		}

		expRes, _ := json.Marshal(resp)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		eCtx := echo.New().NewContext(req, res)

		err := NewJSONResponse(eCtx, statusCode, message, nil)

		assert.Nil(t, err)
		assert.Equal(t, statusCode, res.Code)
		assert.Equal(t, string(expRes)+"\n", res.Body.String())
		assert.NotNil(t, eCtx.Get(consts.ApiResponse))
	})

	t.Run("If the status code 4xx, it will return the error format", func(t *testing.T) {
		message := "Any Message"
		anyErr := errors.New("any error")
		statusCode := http.StatusNotFound

		resp := Error{
			Default: Default{
				Ok:      false,
				Message: message,
			},
			Error: anyErr.Error(),
		}

		expRes, _ := json.Marshal(resp)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		eCtx := echo.New().NewContext(req, res)

		err := NewJSONResponse(eCtx, statusCode, message, anyErr)

		assert.Nil(t, err)
		assert.Equal(t, statusCode, res.Code)
		assert.Equal(t, string(expRes)+"\n", res.Body.String())
		assert.NotNil(t, eCtx.Get(consts.ApiResponse))
	})

	t.Run("If the status code 5xx, it will return the error format", func(t *testing.T) {
		message := "Any Message"
		anyErr := errors.New("any error")
		statusCode := http.StatusInternalServerError

		resp := Error{
			Default: Default{
				Ok:      false,
				Message: message,
			},
			Error: anyErr.Error(),
		}

		expRes, _ := json.Marshal(resp)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		eCtx := echo.New().NewContext(req, res)

		err := NewJSONResponse(eCtx, statusCode, message, anyErr)

		assert.Nil(t, err)
		assert.Equal(t, statusCode, res.Code)
		assert.Equal(t, string(expRes)+"\n", res.Body.String())
		assert.NotNil(t, eCtx.Get(consts.ApiResponse))
	})
}
