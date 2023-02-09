package pondsControllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
	"gorm.io/gorm"
)

// GetPondByID godoc
//
//	@Summary		Get pond data by id
//	@Description	Get pond data by id
//	@Security		JWT
//	@Tags			Ponds
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Pond ID"
//	@Success		200	{object}	generateResp.Success{data=dto.PondResponse,message=string,ok=bool}
//	@Failure		400	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		404	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/ponds/{id} [get]
func (c Controllers) GetByID(ec echo.Context) error {
	id, err := strconv.ParseUint(ec.Param("id"), 10, 64)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	if id == 0 {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, consts.ErrIdRequired.Error(), consts.ErrIdRequired)
	}

	pond, err := c.Usecases.Ponds.GetByID(ec.Request().Context(), id)
	if err != nil {
		return c.buildErrorGetByID(ec, err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewPondResponse(pond))
}

func (c Controllers) buildErrorGetByID(ec echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = http.StatusNotFound
	}

	return generateResp.NewJSONResponse(ec, statusCode, err.Error(), err)
}
