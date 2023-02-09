package pondsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// CreatePond godoc
//
//	@Summary		Create new pond
//	@Description	Create new pond
//	@Security		JWT
//	@Tags			Ponds
//	@Accept			json
//	@Produce		json
//	@Param			farmID	path		uint64					true	"Farm ID"
//	@Param			user	body		dto.CreatePondRequest	true	"Pond Data"
//	@Success		200		{object}	generateResp.Success{data=dto.PondResponse,message=string,ok=bool}
//	@Failure		400		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/farms/{farmID}/ponds [post]
//	@Router			/v1/ponds [post]
func (c Controllers) Create(ec echo.Context) error {
	var req dto.CreatePondRequest
	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	pond, err := c.Usecases.Ponds.Create(ec.Request().Context(), req)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusCreated, "Success", dto.NewPondResponse(pond))
}
