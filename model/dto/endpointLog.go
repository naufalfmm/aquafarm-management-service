package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	RecordRequestLogRequest struct {
		Method string
		Path   string

		RequestID          string
		FullUri            string
		UserAgent          string
		IpAddress          string
		RequestedBy        string
		ResponseStatusCode int
		StartAt            int64
		EndAt              int64

		EndpointID uint64
	}

	EndpointLogReportResponse struct {
		Count                  uint64  `json:"count"`
		UserAgentDistinctCount uint64  `json:"uniqueUserAgent"`
		IpAddressDistinctCount uint64  `json:"uniqueIpAddress"`
		RequestTimeAverage     float64 `json:"requestTimeAverage"`
	}

	EndpointLogReportResponseMap map[string]EndpointLogReportResponse
)

func (req *RecordRequestLogRequest) FromEchoContext(ec echo.Context) error {
	req.Method = ec.Request().Method
	req.Path = cleanPath(ec.Path())

	req.RequestID = ec.Get(consts.XRequestIDHeader).(string)
	req.FullUri = ec.Request().RequestURI
	req.UserAgent = ec.Request().UserAgent()
	req.IpAddress = ec.RealIP()
	req.ResponseStatusCode = ec.Response().Status
	req.StartAt = ec.Get(consts.XRequestStartUnixHeader).(int64)
	req.EndAt = time.Now().UnixMilli()

	req.RequestedBy = consts.SystemCreatedBy
	tokenData, ok := ec.Get(consts.XUserHeader).(token.Data)
	if ok {
		req.RequestedBy = tokenData.CreatedBy()
	}

	return nil
}

func (req RecordRequestLogRequest) ToEndpointLog() dao.EndpointLog {
	splittedUri := strings.Split(req.FullUri, "?")

	query := ""
	if len(splittedUri) > 1 {
		query = splittedUri[1]
	}

	return dao.EndpointLog{
		EndpointID:         req.EndpointID,
		RequestID:          req.RequestID,
		Uri:                splittedUri[0],
		Query:              query,
		UserAgent:          req.UserAgent,
		IpAddress:          req.IpAddress,
		RequestedBy:        &req.RequestedBy,
		ResponseStatusCode: req.ResponseStatusCode,
		StartAt:            req.StartAt,
		EndAt:              &req.EndAt,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		CreatedBy:          consts.SystemCreatedBy,
		UpdatedBy:          consts.SystemCreatedBy,
	}
}

func NewEndpointLogReportResponse(report dao.EndpointLogReport) EndpointLogReportResponse {
	return EndpointLogReportResponse{
		Count:                  report.Count,
		UserAgentDistinctCount: report.UserAgentDistinctCount,
		IpAddressDistinctCount: report.IpAddressDistinctCount,
		RequestTimeAverage:     report.RequestTimeAverage,
	}
}

func NewEndpointLogReportResponseMap(reports dao.EndpointLogReports) EndpointLogReportResponseMap {
	resp := make(EndpointLogReportResponseMap)

	for _, report := range reports {
		resp[fmt.Sprintf("%s %s", report.Endpoint.Method, report.Endpoint.Path)] = NewEndpointLogReportResponse(report)
	}

	return resp
}
