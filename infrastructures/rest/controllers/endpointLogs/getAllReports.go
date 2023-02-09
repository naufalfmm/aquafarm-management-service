package endpointLogsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// GetAllEndpointReports godoc
//	@Summary		Get reports for each endpoint
//	@Description	Get reports for each endpoint
//	@Security		JWT
//	@Tags			Endpoints
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	generateResp.Success{data=dto.EndpointLogReportResponseMap,message=string,ok=bool}
//	@Failure		400	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/endpoints/reports [get]
func (c Controllers) GetAllReports(ec echo.Context) error {
	resp, err := c.Usecases.EndpointLogs.GetAllReports(ec.Request().Context())
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewEndpointLogReportResponseMap(resp))
}
