package farmsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// GetAllPaginatedFarms godoc
//
//	@Summary		Get all paginated farms
//	@Description	Get all farms with pagination and filter
//	@Security		JWT
//	@Tags			Farms
//	@Accept			json
//	@Produce		json
//	@Param			page				query		int		false	"Current Page"
//	@Param			limit				query		int		false	"Count of farms each page"
//	@Param			sorts				query		string	false	"the sorts key of farms sortation"
//	@Param			code				query		string	false	"the farms search by code"
//	@Param			village				query		string	false	"the farms search by village"
//	@Param			district			query		string	false	"the farms search by district"
//	@Param			city				query		string	false	"the farms search by city"
//	@Param			province			query		string	false	"the farms search by province"
//	@Param			postalCode			query		string	false	"the farms search by postal code"
//	@Param			createdDateStart	query		string	false	"the farm created date is filtered greather than or equal created date start"	Format(time)
//	@Param			createdDateEnd		query		string	false	"the farm created date is filtered less than or equal created date end"			Format(time)
//
//	@Success		200					{object}	generateResp.Success{ok=bool,message=string,data=dto.FarmPagingResponse}
//	@Failure		400					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		404					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/farms [get]
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
