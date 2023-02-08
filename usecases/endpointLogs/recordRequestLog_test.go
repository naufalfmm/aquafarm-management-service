package endpointLogsUsecases

import (
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_RecordRequestLog(t *testing.T) {
	var (
		endpoint = dao.Endpoint{
			ID:     1,
			Method: http.MethodGet,
			Path:   "/v1/ponds",
		}

		req = dto.RecordRequestLogRequest{
			Method:     endpoint.Method,
			Path:       endpoint.Path,
			EndpointID: endpoint.ID,
		}

		endpointLog = dao.EndpointLog{
			ID:         1,
			EndpointID: endpoint.ID,
		}
	)

	t.Run("If no error, it will return the recorded log", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.endpointRepositories.EXPECT().GetByMethodPath(mock.ctx, req.Method, req.Path).Return(endpoint, nil)
		mock.endpointLogRepositories.EXPECT().Create(mock.ctx, req.ToEndpointLog()).Return(endpointLog, nil)

		res, err := mock.usecases.RecordRequestLog(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, endpointLog, res)
	})

	t.Run("If get endpoint by method and path return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.endpointRepositories.EXPECT().GetByMethodPath(mock.ctx, req.Method, req.Path).Return(dao.Endpoint{}, errAny)

		res, err := mock.usecases.RecordRequestLog(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.EndpointLog{}, res)
	})
}
