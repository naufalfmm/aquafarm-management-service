package farmsControllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
	"gorm.io/gorm"
)

// DeleteFarmByID godoc
//	@Summary		Delete farm by id
//	@Description	Delete farm by id and ponds connected to farm id
//	@Security		JWT
//	@Tags			Farms
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Farm ID"
//	@Success		200	{object}	generateResp.Success{ok=bool,message=string}
//	@Failure		400	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		404	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500	{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/farms/{id} [delete]
func (c Controllers) DeleteByID(ec echo.Context) error {
	id, err := strconv.ParseUint(ec.Param("id"), 10, 64)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	if id == 0 {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, consts.ErrIdRequired.Error(), consts.ErrIdRequired)
	}

	loginData := ec.Get(consts.XUserHeader).(token.Data)

	if err := c.Usecases.Farms.DeleteByID(ec.Request().Context(), id, loginData.CreatedBy()); err != nil {
		return c.buildErrorDeleteByID(ec, err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", nil)
}

func (c Controllers) buildErrorDeleteByID(ec echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = http.StatusNotFound
	}

	return generateResp.NewJSONResponse(ec, statusCode, err.Error(), err)
}
