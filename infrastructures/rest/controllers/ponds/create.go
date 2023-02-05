package pondsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

func (c Controllers) Create(ec echo.Context) error {
	var req dto.CreatePondRequest
	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	pond, err := c.Usecases.Ponds.Create(ec.Request().Context(), req)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewPondResponse(pond))
}
