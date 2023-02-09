package hooks

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
)

func (h Hooks) BulkUpsertEndpoints(ec echo.Context) error {
	var req dto.BulkUpsertEndpointsRequest
	if err := req.NewFromEcho(*ec.Echo()); err != nil {
		return err
	}

	if err := h.Usecases.Endpoints.BulkUpsertEndpoints(context.Background(), req); err != nil {
		h.Resources.Logger.Error(context.Background(),
			"error when bulk upserting endpoints",
			zapLog.SetAttribute("req", req),
		)
		return err
	}

	return nil
}
