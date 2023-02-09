package farmsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// UpsertFarm godoc
//
//	@Summary		Update or insert farm
//	@Description	Update the farm if farm is exist or insert if farm is not exist
//	@Security		JWT
//	@Tags			Farms
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string						true	"Farm Code"
//	@Param			user	body		dto.UpsertFarmBodyRequest	true	"Farm Data"
//	@Success		200		{object}	generateResp.Success{data=dto.FarmResponse,message=string,ok=bool}
//	@Failure		400		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/farms/{code} [put]
func (c Controllers) Upsert(ec echo.Context) error {
	var req dto.UpsertFarmRequest
	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	farm, err := c.Usecases.Farms.Upsert(ec.Request().Context(), req)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewFarmResponse(farm))
}
