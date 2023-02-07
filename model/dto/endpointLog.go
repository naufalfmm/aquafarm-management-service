package dto

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	RequestStartLogRequest struct {
		Method string
		Path   string

		RequestID   string
		FullUri     string
		UserAgent   string
		IpAddress   string
		RequestedBy string

		EndpointID uint64
	}
)

func (req *RequestStartLogRequest) FromEchoContext(ec echo.Context) error {
	req.Method = ec.Request().Method
	req.Path = cleanPath(ec.Path())

	req.RequestID = ec.Get(consts.XRequestIDHeader).(string)
	req.FullUri = ec.Request().RequestURI
	req.UserAgent = ec.Request().UserAgent()
	req.IpAddress = ec.RealIP()
	req.RequestedBy = ec.Get(consts.XUserHeader).(token.Data).CreatedBy()

	return nil
}

func (req RequestStartLogRequest) ToEndpointLog() dao.EndpointLog {
	splittedUri := strings.Split(req.FullUri, "?")

	query := ""
	if len(splittedUri) > 1 {
		query = splittedUri[1]
	}

	return dao.EndpointLog{
		EndpointID:  req.EndpointID,
		RequestID:   req.RequestID,
		Uri:         splittedUri[0],
		Query:       query,
		UserAgent:   req.UserAgent,
		IpAddress:   req.IpAddress,
		RequestedBy: &req.RequestedBy,
		StartAt:     time.Now().UnixMilli(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   consts.SystemCreatedBy,
		UpdatedBy:   consts.SystemCreatedBy,
	}
}
