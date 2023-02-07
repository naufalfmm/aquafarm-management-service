package endpointLogsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

func (c Controllers) GetAllReports(ec echo.Context) error {
	resp, err := c.Usecases.EndpointLogs.GetAllReports(ec.Request().Context())
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewEndpointLogReportResponseMap(resp))
}
