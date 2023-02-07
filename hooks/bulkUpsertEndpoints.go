package hooks

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
)

func (h Hooks) BulkUpsertEndpoints(ec echo.Context) error {
	var req dto.BulkUpsertEndpointsRequest
	if err := req.NewBulkUpsertEndpointsRequestFromEcho(*ec.Echo()); err != nil {
		return err
	}

	return h.Usecases.Endpoints.BulkUpsertEndpoints(context.Background(), req)
}
