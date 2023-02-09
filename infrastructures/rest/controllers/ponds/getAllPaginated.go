package pondsControllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

// GetAllPaginatedPonds godoc
//
//	@Summary		Get all paginated ponds
//	@Description	Get all ponds with pagination and filter
//	@Security		JWT
//	@Tags			Ponds
//	@Accept			json
//	@Produce		json
//
//	@Param			page				query		int		false	"Current Page"
//	@Param			limit				query		int		false	"Count of ponds each page"
//	@Param			sorts				query		string	false	"the sorts key of ponds sortation"
//	@Param			code				query		string	false	"the ponds search by code"
//	@Param			volumeStart			query		decimal	false	"the pond volume is filtered greather than or equal volume start"
//	@Param			volumeEnd			query		decimal	false	"the pond volume is filtered less than or equal volume end"
//	@Param			areaStart			query		decimal	false	"the pond area is filtered greather than or equal area start"
//	@Param			areaEnd				query		decimal	false	"the pond area is filtered less than or equal area end"
//	@Param			createdDateStart	query		string	false	"the pond created date is filtered greather than or equal created date start"	Format(time)
//	@Param			createdDateEnd		query		string	false	"the pond created date is filtered less than or equal created date end"			Format(time)
//
//	@Success		200					{object}	generateResp.Success{ok=bool,message=string,data=dto.PondPagingResponse}
//	@Failure		400					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		404					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Failure		500					{object}	generateResp.Error{ok=bool,message=string,error=error}
//	@Router			/v1/ponds [get]
func (c Controllers) GetAllPaginated(ec echo.Context) error {
	var req dto.PondPagingRequest

	if err := req.FromEchoContext(ec); err != nil {
		return generateResp.NewJSONResponse(ec, http.StatusBadRequest, err.Error(), err)
	}

	resp, err := c.Usecases.Ponds.GetAllPaginated(ec.Request().Context(), req)
	if err != nil {
		return c.buildErrorGetAllPaginated(ec, err)
	}

	return generateResp.NewJSONResponse(ec, http.StatusOK, "Success", dto.NewPondPagingResponse(resp))
}

func (c Controllers) buildErrorGetAllPaginated(ec echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	switch err {
	case consts.ErrRecordNotFound:
		statusCode = http.StatusNotFound
	}

	return generateResp.NewJSONResponse(ec, statusCode, err.Error(), err)
}
