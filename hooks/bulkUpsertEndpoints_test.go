package hooks

import (
	"context"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"github.com/stretchr/testify/assert"
)

func Test_hook_BulkUpsertEndpoints(t *testing.T) {
	var (
		req = dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
			},
		}
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetEchoRoute(req.Endpoints[0].Method, req.Endpoints[0].Path, func(c echo.Context) error { return nil })

		mock.endpointUsecases.EXPECT().BulkUpsertEndpoints(context.Background(), req).Return(nil)

		err := mock.hooks.BulkUpsertEndpoints(mock.eCtx)

		assert.Nil(t, err)
	})

	t.Run("If bulk upsert error, it will return error and log it", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetEchoRoute(req.Endpoints[0].Method, req.Endpoints[0].Path, func(c echo.Context) error { return nil })

		mock.endpointUsecases.EXPECT().BulkUpsertEndpoints(context.Background(), req).Return(errAny)
		mock.logger.EXPECT().Error(context.Background(),
			"error when bulk upserting endpoints",
			zapLog.SetAttribute("req", req),
		)

		err := mock.hooks.BulkUpsertEndpoints(mock.eCtx)

		assert.Equal(t, errAny, err)
	})
}
