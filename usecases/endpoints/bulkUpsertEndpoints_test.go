package endpointsUsecases

import (
	"net/http"
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_BulkUpsertEndpoints(t *testing.T) {
	t.Run("If there is unused endpoint in echo and unfound endpoint in db, the unused will be deleted and unfound will be added", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		endpoints := dao.Endpoints{
			{
				ID:     1,
				Method: http.MethodGet,
				Path:   "/v1/ponds",
			},
			{
				ID:     2,
				Method: http.MethodGet,
				Path:   "/v1/farms",
			},
		}

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
				{
					Method: http.MethodPost,
					Path:   "/v1/ponds",
				},
				{
					Method: http.MethodPost,
					Path:   "/v1/farms",
				},
			},
		}

		addedEndpoints := dao.Endpoints{
			req.Endpoints[1].ToEndpoint(),
			req.Endpoints[2].ToEndpoint(),
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(endpoints, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)

		mock.endpointRepositories.EXPECT().BulkDeleteByIDs(mock.ctx, []uint64{2}).Return(nil)
		mock.endpointRepositories.EXPECT().BulkCreate(mock.ctx, addedEndpoints).Return(addedEndpoints, nil)

		mock.orm.EXPECT().Commit()

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Nil(t, err)
	})

	t.Run("If there is unused endpoint in echo, the unused will be deleted and no endpoint will be added", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		endpoints := dao.Endpoints{
			{
				ID:     1,
				Method: http.MethodGet,
				Path:   "/v1/ponds",
			},
			{
				ID:     2,
				Method: http.MethodGet,
				Path:   "/v1/farms",
			},
		}

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
			},
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(endpoints, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)

		mock.endpointRepositories.EXPECT().BulkDeleteByIDs(mock.ctx, []uint64{2}).Return(nil)

		mock.orm.EXPECT().Commit()

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Nil(t, err)
	})

	t.Run("If there is unfound endpoint in db, the unfound will be added and no endpoint will be deleted", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		endpoints := dao.Endpoints{
			{
				ID:     1,
				Method: http.MethodGet,
				Path:   "/v1/ponds",
			},
			{
				ID:     2,
				Method: http.MethodGet,
				Path:   "/v1/farms",
			},
		}

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
				{
					Method: http.MethodGet,
					Path:   "/v1/farms",
				},
				{
					Method: http.MethodPost,
					Path:   "/v1/farms",
				},
			},
		}

		addedEndpoints := dao.Endpoints{
			req.Endpoints[2].ToEndpoint(),
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(endpoints, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)

		mock.endpointRepositories.EXPECT().BulkCreate(mock.ctx, addedEndpoints).Return(addedEndpoints, nil)

		mock.orm.EXPECT().Commit()

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Nil(t, err)
	})

	t.Run("If bulk create endpoints return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		endpoints := dao.Endpoints{
			{
				ID:     1,
				Method: http.MethodGet,
				Path:   "/v1/ponds",
			},
			{
				ID:     2,
				Method: http.MethodGet,
				Path:   "/v1/farms",
			},
		}

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
				{
					Method: http.MethodGet,
					Path:   "/v1/farms",
				},
				{
					Method: http.MethodPost,
					Path:   "/v1/farms",
				},
			},
		}

		addedEndpoints := dao.Endpoints{
			req.Endpoints[2].ToEndpoint(),
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(endpoints, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)
		mock.orm.EXPECT().Rollback()

		mock.endpointRepositories.EXPECT().BulkCreate(mock.ctx, addedEndpoints).Return(nil, errAny)

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Equal(t, errAny, err)
	})

	t.Run("If bulk delete endpoints by ids return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		endpoints := dao.Endpoints{
			{
				ID:     1,
				Method: http.MethodGet,
				Path:   "/v1/ponds",
			},
			{
				ID:     2,
				Method: http.MethodGet,
				Path:   "/v1/farms",
			},
		}

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
			},
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(endpoints, nil)

		mock.orm.EXPECT().Begin().Return(mock.orm)
		mock.orm.EXPECT().Rollback()

		mock.endpointRepositories.EXPECT().BulkDeleteByIDs(mock.ctx, []uint64{2}).Return(errAny)

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Equal(t, errAny, err)
	})

	t.Run("If get all endpoints return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.BulkUpsertEndpointsRequest{
			Endpoints: dto.UpsertEndpointRequests{
				{
					Method: http.MethodGet,
					Path:   "/v1/ponds",
				},
			},
		}

		mock.endpointRepositories.EXPECT().GetAll(mock.ctx).Return(nil, errAny)

		err := mock.usecases.BulkUpsertEndpoints(mock.ctx, req)

		assert.Equal(t, errAny, err)
	})
}
