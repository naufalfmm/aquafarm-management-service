package farmsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// CreateFarm godoc
//	@Summary		Create new farm
//	@Description	Create new farm
//	@Security		JWT
//	@Tags			Farms
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.CreateFarmRequest	true	"Farm Data"
//	@Success		200		{object}	generateResp.Success{data=dto.FarmResponse,message=string,ok=bool}
//	@Failure		400		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/farms [post]
func (c Controllers) Create(ec echo.Context) error {
	var req dto.CreateFarmRequest
	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	farm, err := c.Usecases.Farms.Create(ec.Request().Context(), req)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusCreated, "Success", dto.NewFarmResponse(farm))
}
