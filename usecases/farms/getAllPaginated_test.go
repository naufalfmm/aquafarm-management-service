package farmsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_GetAllPaginated(t *testing.T) {
	var (
		req = dto.FarmPagingRequest{}

		pagingResp = dao.FarmsPagingResponse{
			Items: dao.Farms{
				{ID: 1},
			},
		}
	)

	t.Run("If no error and the items is not empty, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetAllPaginated(mock.ctx, req).Return(pagingResp, nil)

		res, err := mock.usecases.GetAllPaginated(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, pagingResp, res)
	})

	t.Run("If the items is empty, it will return record not found", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		pagingResp = dao.FarmsPagingResponse{
			Items: dao.Farms{},
		}

		mock.farmRepositories.EXPECT().GetAllPaginated(mock.ctx, req).Return(pagingResp, nil)

		res, err := mock.usecases.GetAllPaginated(mock.ctx, req)

		assert.Equal(t, consts.ErrRecordNotFound, err)
		assert.Equal(t, dao.FarmsPagingResponse{}, res)
	})

	t.Run("If get all paginated farms error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		pagingResp = dao.FarmsPagingResponse{
			Items: dao.Farms{},
		}

		mock.farmRepositories.EXPECT().GetAllPaginated(mock.ctx, req).Return(dao.FarmsPagingResponse{}, errAny)

		res, err := mock.usecases.GetAllPaginated(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.FarmsPagingResponse{}, res)
	})
}
