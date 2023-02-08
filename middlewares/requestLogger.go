package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (m middlewares) RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			m.logRequest(c)

			err := next(c)

			m.logResponse(c, err)

			return err
		}
	}
}

func (m middlewares) logRequest(c echo.Context) {
	c.Set(consts.XRequestIDHeader, uuid.New().String())
	c.Set(consts.XRequestStartUnixHeader, time.Now().UnixMilli())

	var mapRequest map[string]interface{}
	if c.Request().Body != nil {
		requestByte, _ := ioutil.ReadAll(c.Request().Body)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(requestByte)) // Reset
		_ = json.Unmarshal(requestByte, &mapRequest)
	}

	requestMsg := fmt.Sprintf("[api-log] - [REQUEST] - [%s] %s",
		c.Request().Method,
		c.Request().URL.Path,
	)

	if c.Request().URL.RawQuery != "" {
		requestMsg += fmt.Sprintf("?%s", c.Request().URL.RawQuery)
	}

	m.Resources.Logger.Info(c.Request().Context(), requestMsg,
		zapLog.SetAttribute("payload", mapRequest),
		zapLog.SetAttribute("headers", c.Request().Header),
	)
}

func (m middlewares) recordRequestLog(c echo.Context) {
	_, ok := c.Get(consts.XRequestIDHeader).(string)
	if !ok {
		return
	}

	req := dto.RecordRequestLogRequest{}
	if err := req.FromEchoContext(c); err != nil {
		return
	}

	go m.Usecases.EndpointLogs.RecordRequestLog(context.Background(), req)
}

func (m middlewares) logResponse(c echo.Context, err error) {
	responseMsg := fmt.Sprintf("[api-log] - [RESPONSE] - [%d] [%s] %s",
		c.Response().Status,
		c.Request().Method,
		c.Request().URL.Path)
	if c.Request().URL.RawQuery != "" {
		responseMsg += fmt.Sprintf("?%s", c.Request().URL.RawQuery)
	}

	if err != nil {
		m.Resources.Logger.Error(c.Request().Context(), responseMsg, zapLog.SetAttribute("error", err))
	} else {
		m.Resources.Logger.Info(c.Request().Context(), responseMsg, zapLog.SetAttribute("response", c.Get(consts.ApiResponse)))
	}

	m.recordRequestLog(c)
}
