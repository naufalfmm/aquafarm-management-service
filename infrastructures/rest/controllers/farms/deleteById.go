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

func (c Controllers) DeleteByID(ec echo.Context) error {
	id, err := strconv.ParseUint(ec.Param("id"), 10, 64)
	if err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	if id == 0 {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, consts.ErrIdRequired.Error(), consts.ErrIdRequired)
	}

	loginData := ec.Get("x-user").(token.Data)

	if err := c.Usecases.Farms.DeleteByID(ec.Request().Context(), id, loginData); err != nil {
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
