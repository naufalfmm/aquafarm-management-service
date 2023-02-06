package farmsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

func (c Controllers) GetAllPaginated(ec echo.Context) error {
	var req dto.FarmPagingRequest

	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	resp, err := c.Usecases.Farms.GetAllPaginated(ec.Request().Context(), req)
	if err != nil {
		return c.buildErrorGetAllPaginated(ec, err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewFarmPagingResponse(resp))
}

func (c Controllers) buildErrorGetAllPaginated(ec echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	switch err {
	case consts.ErrRecordNotFound:
		statusCode = http.StatusNotFound
	}

	return generateResp.NewJSONResponse(ec, statusCode, err.Error(), err)
}
