package pondsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// UpsertPond godoc
//
//	@Summary		Update or insert pond
//	@Description	Update the pond if pond is exist or insert if pond is not exist
//	@Security		JWT
//	@Tags			Ponds
//	@Accept			json
//	@Produce		json
//
//	@Param			farmID	path		int						true	"Farm ID"
//	@Param			code	path		string					true	"Pond Code"
//	@Param			user	body		dto.UpsertPondRequest	true	"Pond Data"
//
//	@Success		200		{object}	generateResp.Success{data=dto.PondResponse,message=string,ok=bool}
//	@Failure		400		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500		{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/ponds/{code} [put]
//	@Router			/v1/farms/{farmID}/ponds/{code} [put]
func (c Controllers) Upsert(ec echo.Context) error {
	var req dto.UpsertPondRequest
	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	pond, err := c.Usecases.Ponds.Upsert(ec.Request().Context(), req)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusInternalServerError, err.Error(), err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewPondResponse(pond))
}
